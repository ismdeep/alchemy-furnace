package config

import (
	"context"
	"github.com/ismdeep/log"
	etcdClientV3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
	"time"
)

func load() {
	log.Info("init", log.String("info", "started to load config"))
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

	resp, err := cli.Get(context.Background(), AlchemyFurnaceConfigName)
	if err != nil {
		return
	}

	for _, kv := range resp.Kvs {
		if string(kv.Key) == AlchemyFurnaceConfigName {
			c := &config{}
			_ = yaml.Unmarshal(kv.Value, c)
			Config = c
			log.Info("init", log.String("info", "config loaded"))
		}
	}

}
