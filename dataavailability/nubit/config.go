package nubit

import (
	"encoding/json"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"io"
	"os"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/config/types"
)

// NubitNamespaceBytesLength is the fixed-size bytes array.
const NubitNamespaceBytesLength = 58

// NubitMinCommitTime is the minimum commit time interval between blob submissions to NubitDA.
const NubitMinCommitTime time.Duration = 12 * time.Second

// Config is the NubitDA backend configurations
type Config struct {
	NubitRpcURL             string         `mapstructure:"NubitRpcURL"`
	NubitAuthKey            string         `mapstructure:"NubitAuthKey"`
	NubitNamespace          string         `mapstructure:"NubitNamespace"`
	NubitGetProofMaxRetry   uint64         `mapstructure:"NubitGetProofMaxRetry"`
	NubitGetProofWaitPeriod types.Duration `mapstructure:"NubitGetProofWaitPeriod"`
}

func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	log.Infof("⚙️     Nubit byteValue : %s ", string(byteValue))
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return err
	}
	log.Infof("⚙️     Nubit GetConfig : %#v ", c)
	return nil
}
