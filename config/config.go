package config

import (
	"context"
	"fmt"
	etcdClientV3 "go.etcd.io/etcd/client/v3"
	"os"
	"time"
)

// ALCHEMY_FURNACE_ROOT
// ALCHEMY_FURNACE_ETCD

const AlchemyFurnaceConfigName = "alchemy-furnace-config"

var WorkDir string     // working directory
var etcdAddress string // etcd address

type config struct {
	Bind string `yaml:"bind"`
	DSN  string `yaml:"dsn"`
}

var Config *config

func init() {
	// 1. 获取工作目录
	WorkDir, _ = os.Getwd()
	if os.Getenv("ALCHEMY_FURNACE_ROOT") != "" {
		WorkDir = os.Getenv("ALCHEMY_FURNACE_ROOT")
	}

	// 2. 获取ETCD地址
	etcdAddress = os.Getenv("ALCHEMY_FURNACE_ETCD")
	if etcdAddress == "" {
		fmt.Println("ALCHEMY_FURNACE_ETCD is empty")
		os.Exit(1)
	}

	// 3. 加载配置
	Config = &config{}
	load()

	// 4. 监听ETCD配置变化
	go func() {
		cli, err := etcdClientV3.New(etcdClientV3.Config{
			Endpoints:   []string{etcdAddress},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			panic(err)
		}

		watcher := cli.Watch(context.Background(), AlchemyFurnaceConfigName)
		for {
			e := <-watcher
			for _, event := range e.Events {
				if string(event.Kv.Key) == AlchemyFurnaceConfigName {
					load()
					for _, w := range Watchers {
						w <- true
					}
				}
			}
		}
	}()
}
