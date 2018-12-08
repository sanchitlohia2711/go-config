package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	//ErrFailedToDecodeConfigurationFile : err failed decoding configuration file
	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v"

	//ErrFailedToOpenConfigurationFile : info using default configuration file path
	ErrFailedToOpenConfigurationFile = "Failed to open configuration file: %v"
)

var (
	configurationFilePath string
)

//Initialize :the config
func Initialize(configurationFileAbsPath string, configurationData interface{}) (err error) {
	configFile, err := os.Open(configurationFileAbsPath)
	defer configFile.Close()
	if err != nil {
		err = fmt.Errorf(ErrFailedToOpenConfigurationFile, err)
		return
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configurationData)
	if err != nil {
		err = fmt.Errorf(ErrFailedToDecodeConfigurationFile, err)
	}
	return
}
