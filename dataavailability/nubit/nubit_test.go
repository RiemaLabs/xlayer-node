package nubit

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStoreDetailsOnChain(t *testing.T) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	nubit, err := NewNubitDABackendTest("http://127.0.0.1:26658", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.DAMv0s7915Ahx-kDFSzDT1ATz4Q9WwktWcHmjp7_99Q", pk)
	if err != nil {
		t.Fatal(err)
	}
	txs := []byte("test txs")
	returnData, _, _, err := nubit.PostSequence(context.TODO(), [][]byte{txs})
	require.NoError(t, err)
	_, err = nubit.GetSequence(context.TODO(), []common.Hash{}, returnData)
	require.NoError(t, err)
}
