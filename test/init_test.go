package test

import (
	"github.com/ismdeep/alchemy-furnace/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.Config.Bind = "0.0.0.0:8000"
	config.Config.DSN = "root:liandanlu123456@tcp(127.0.0.1:10006)/alchemy_furnace?parseTime=true&loc=Local&charset=utf8mb4,utf8"
	config.Config.JWT = "xMTMgCwGSVw97AxnM6aaAbSB7GKrF9Jq"
	config.Dump()
	m.Run()
}
