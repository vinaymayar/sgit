package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/vinaymayar/sgit/utils"
)

type Config struct {
	ProjectType string
}

const (
	configFileName = "sgit.config"
)

func (config Config) Write() error {
	out, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, out, 0666)
}

func Configure(args []string) {
	config, err := GetConfig()
	if err != nil && !os.IsNotExist(err) {
		log.Print(err)
	}

	if len(args) > 1 {
		config.ProjectType = args[1]
	}

	err = config.Write()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() (Config, error) {
	config := new(Config)
	configPath, err := getConfigPath()
	if err != nil {
		return *config, err
	}

	configBuf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return *config, err
	}

	if err := json.Unmarshal(configBuf, config); err != nil {
		return *config, err
	}
	return *config, err
}

func getConfigPath() (string, error) {
	sgitDir, err := utils.GetSgitRootDir()
	if err != nil {
		return sgitDir, err
	}

	configPath := filepath.Join(sgitDir, configFileName)
	return configPath, err
}
