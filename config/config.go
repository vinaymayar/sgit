package config

import (
	"encoding/json"
	"flag"
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
	unset          = "unset"
	configFileName = "sgit.config"
	configCmd      = "configure"
)

var projectType string

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

func IsConfigCmd(cmd string) bool {
	return cmd == configCmd
}

func Configure() {
	config, err := GetConfig()
	if err != nil && !os.IsNotExist(err) {
		log.Print(err)
	}

	if isSet(projectType) {
		config.ProjectType = projectType
	}

	err = config.Write()
	if err != nil {
		log.Fatal(err)
	}
}

func InitFlags() {
	flag.StringVar(&projectType, "p", unset, "project type.  Run `sgit` for more information.")
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

func isSet(str string) bool {
	return str != unset
}
