package nubit

import (
	"encoding/json"
	"io"
	"os"

	"github.com/0xPolygonHermez/zkevm-node/log"
)

type Config struct {
	RpcURL    string `json:"rpcURL"`
	Namespace string `json:"modularAppName"`
	AuthKey   string `json:"authKey"`
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
