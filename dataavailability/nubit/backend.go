package nubit

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	daTypes "github.com/0xPolygon/cdk-data-availability/types"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/polygondatacommittee"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"time"
)

// // DABackender is an interface for components that store and retrieve batch data
// type DABackender interface {
// 	SequenceRetriever
// 	SequenceSender
// 	// Init initializes the DABackend
// 	Init() error
// }

// // SequenceSender is used to send provided sequence of batches
// type SequenceSender interface {
// 	// PostSequence sends the sequence data to the data availability backend, and returns the dataAvailabilityMessage
// 	// as expected by the contract
// 	PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error)
// }

// // SequenceRetriever is used to retrieve batch data
// type SequenceRetriever interface {
// 	// GetSequence retrieves the sequence data from the data availability backend
// 	GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error)
// }

type NubitDABackend struct {
	commitTime            time.Time
	config                *Config
	attestationContract   *polygondatacommittee.Polygondatacommittee
	ns                    da.Namespace
	privKey               *ecdsa.PrivateKey
	client                da.DA
	dataCommitteeContract *polygondatacommittee.Polygondatacommittee
}

func NewNubitDABackend(l1RPCURL string, dataCommitteeAddr common.Address, privKey *ecdsa.PrivateKey) (*NubitDABackend, error) {
	var config Config
	err := config.GetConfig("/app/nubit-config.json")
	if err != nil {
		log.Fatalf("cannot get config:%w", err)
	}

	ethClient, err := ethclient.Dial(l1RPCURL)
	if err != nil {
		log.Errorf("error connecting to %s: %+v", l1RPCURL, err)
		return nil, err
	}

	log.Infof("⚙️     Nubit config : %#v ", config)

	attestationContract, err := polygondatacommittee.NewPolygondatacommittee(dataCommitteeAddr, ethClient)
	if err != nil {
		return nil, err
	}
	cn, err := proxy.NewClient(config.RpcURL, config.AuthKey)
	if err != nil {
		return nil, err
	}
	name, err := hex.DecodeString("00000000000000000000000000000000000000000000706f6c79676f6e")
	dataCommittee, err := polygondatacommittee.NewPolygondatacommittee(dataCommitteeAddr, ethClient)

	log.Infof("⚙️     Nubit Namespace : %s ", string(name))
	return &NubitDABackend{
		dataCommitteeContract: dataCommittee,
		config:                &config,
		attestationContract:   attestationContract,
		privKey:               privKey,
		ns:                    name,
		client:                cn,
		commitTime:            time.Now(),
	}, nil
}

func NewNubitDABackendTest(url string, authKey string, pk *ecdsa.PrivateKey) (*NubitDABackend, error) {
	cn, err := proxy.NewClient(url, authKey)
	if err != nil || cn == nil {
		return nil, err
	}

	name, err := hex.DecodeString("00000000000000000000000000000000000000000000706f6c79676f6e")
	if err != nil {
		return nil, err
	}

	log.Infof("⚙️     Nubit Namespace : %s ", string(name))
	return &NubitDABackend{
		ns:         name,
		client:     cn,
		privKey:    pk,
		commitTime: time.Now(),
	}, nil
}

func (a *NubitDABackend) Init() error {
	return nil
}

// PostSequence sends the sequence data to the data availability backend, and returns the dataAvailabilityMessage
// as expected by the contract

var BatchsDataCache = [][]byte{}
var BatchsSize = 0

func (a *NubitDABackend) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, []byte, error) {
	encodedData, err := MarshalBatchData(batchesData)
	if err != nil {
		log.Errorf("🏆    NubitDABackend.MarshalBatchData:%s", err)
		return encodedData, nil, err
	}

	BatchsDataCache = append(BatchsDataCache, encodedData)
	BatchsSize += len(encodedData)
	if BatchsSize < 100*1024 {
		log.Infof("🏆  Nubit BatchsDataCache:%+v", len(encodedData))
		return nil, nil, nil
	}
	if time.Since(a.commitTime) < 12*time.Second {
		time.Sleep(time.Since(a.commitTime))
	}

	BatchsData, err := MarshalBatchData(BatchsDataCache)
	if err != nil {
		log.Errorf("🏆    NubitDABackend.MarshalBatchData:%s", err)
		return nil, nil, err
	}
	id, err := a.client.Submit(ctx, [][]byte{BatchsData}, -1, a.ns)
	if err != nil {
		log.Errorf("🏆    NubitDABackend.Submit:%s", err)
		return nil, nil, err
	}

	log.Infof("🏆  Nubit Data submitted by sequencer: %d bytes against namespace %v sent with id %#x", len(BatchsDataCache), a.ns, id)
	a.commitTime = time.Now()
	BatchsDataCache = [][]byte{}
	BatchsSize = 0
	// todo: May be need to sleep
	//dataProof, err := a.client.Blob.GetProof(ctx, uint64(blockNumber), a.ns.Bytes(), body.Commitment)
	//if err != nil {
	//	log.Errorf("🏆    NubitDABackend.GetProof:%s", err)
	//	return nil, err
	//}
	//
	//log.Infof("🏆   Nubit received data proof:%+v", dataProof)

	var batchDAData BatchDAData
	batchDAData.ID = id
	log.Infof("🏆  Nubit prepared DA data:%+v", batchDAData)

	// todo: use bridge API data
	returnData, err := batchDAData.Encode()
	if err != nil {
		return nil, nil, fmt.Errorf("🏆  Nubit cannot encode batch data:%w", err)
	}

	sequence := daTypes.Sequence{}
	for _, seq := range batchesData {
		sequence = append(sequence, seq)
	}
	signedSequence, err := sequence.Sign(a.privKey) //todo
	sequence.HashToSign()

	Signature := append(sequence.HashToSign(), signedSequence.Signature...)

	return returnData, Signature, nil
}

func (a *NubitDABackend) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	batchDAData := BatchDAData{}
	err := batchDAData.Decode(dataAvailabilityMessage)
	if err != nil {
		log.Errorf("🏆    NubitDABackend.GetSequence.Decode:%s", err)
		return nil, err
	}
	log.Infof("🏆     Nubit GetSequence batchDAData:%+v", batchDAData)
	blob, err := a.client.Get(context.TODO(), batchDAData.ID, a.ns)
	if err != nil {
		log.Errorf("🏆    NubitDABackend.GetSequence.Blob.Get:%s", err)
		return nil, err
	}
	log.Infof("🏆     Nubit GetSequence blob.data:%+v", len(blob))
	byteBlob := make([][]byte, len(blob))
	for _, b := range blob {
		byteBlob = append(byteBlob, b)
	}
	return byteBlob, nil
}

// DataCommitteeMember represents a member of the Data Committee
type DataCommitteeMember struct {
	Addr common.Address
	URL  string
}

// DataCommittee represents a specific committee
type DataCommittee struct {
	AddressesHash      common.Hash
	Members            []DataCommitteeMember
	RequiredSignatures uint64
}
