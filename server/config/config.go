package config

import (
	"fmt"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// WorkDir working directory ALCHEMY_FURNACE_ROOT
var WorkDir string

type config struct {
	Bind  string
	JWT   string
	Token string // X-Token
	Auth  struct {
		Username string
		Password string
	}
	WeCom string
}

var ROOT config

func init() {
	// 1. 获取工作目录
	WorkDir = os.Getenv("ALCHEMY_FURNACE_ROOT")
	if WorkDir == "" {
		fmt.Println("Please set ALCHEMY_FURNACE_ROOT")
		os.Exit(1)
	}
	if err := os.MkdirAll(fmt.Sprintf("%v/data", WorkDir), 0777); err != nil {
		panic(err)
	}
	log.Info("init", log.String("WorkDir", WorkDir))

	raw, err := ioutil.ReadFile(fmt.Sprintf("%v/config.yaml", WorkDir))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(raw, &ROOT); err != nil {
		panic(err)
	}

	jwt.Init(&jwt.Config{
		Key:    ROOT.JWT,
		Expire: "72h",
	})
}
