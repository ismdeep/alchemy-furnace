package config

import (
	"context"
	"github.com/ismdeep/log"
	etcdClientV3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"time"
)

func Dump() {
	if Config == nil {
		return
	}

	bytes, err := yaml.Marshal(Config)
	if err != nil {
		panic(err)
	}

	cli, err := etcdClientV3.New(etcdClientV3.Config{
		Endpoints:   []string{etcdAddress},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			log.Error("config", log.FieldErr(err))
		}
	}()

	if _, err := cli.Put(context.Background(), AlchemyFurnaceConfigName, string(bytes)); err != nil {
		log.Error("config", log.FieldErr(err))
	}
}
