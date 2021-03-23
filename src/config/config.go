package conf

import (
	"encoding/json"
	"errors"
	"os"
)

type EmailQuotaConf struct {
	APIKey *string `json:"ApiKey"`
}

var Configuration EmailQuotaConf

func SetupConfiguration(confPath string) error {
	if confFile, err := os.Open(confPath); err != nil {
		return err
	} else {
		decoder := json.NewDecoder(confFile)
		if err := decoder.Decode(&Configuration); err != nil {
			return err
		}
	}
	if Configuration.APIKey == nil {
		return errors.New("APIKey not set")
	}
	return nil
}

func GetConfiguration() EmailQuotaConf {
	return Configuration
}
