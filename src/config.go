package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//connection json wrapper
type connConfig struct {
	Conf Config `json:"Connection"`
}

//Stores config info
type Config struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetConfig() (Config, error) {
	//Create a connection wrapper
	connConfig := connConfig{}

	//reading data from config
	f, err := os.Open("./config.json")
	if err != nil {
		return Config{}, err
	}
	configData, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(configData, &connConfig)
	if err != nil {
		return Config{}, err
	}
	return connConfig.Conf, nil
}
