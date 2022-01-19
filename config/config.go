package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type config struct {
	DB struct {
		DSN string `yaml:"dsn"`
	} `yaml:"db"`
	Tasks []struct {
		ID   string `yaml:"id"`
		Name string `yaml:"name"`
		Bash string `yaml:"bash"`
		Cron string `yaml:"cron"`
	} `yaml:"tasks"`
}

var Config *config
var WorkDir string

func init() {
	WorkDir, _ = os.Getwd()
	if os.Getenv("ALCHEMY_FURNACE_ROOT") != "" {
		WorkDir = os.Getenv("ALCHEMY_FURNACE_ROOT")
	}

	bytes, err := ioutil.ReadFile(fmt.Sprintf("%v/config.yaml", WorkDir))
	if err != nil {
		panic(err)
	}

	Config = &config{}
	if err := yaml.Unmarshal(bytes, Config); err != nil {
		panic(err)
	}

	for _, task := range Config.Tasks {
		fmt.Println(task.Cron, task.ID, task.Name)
	}
}
