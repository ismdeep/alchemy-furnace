package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"testing"
)

var testTaskID uint

func TestMain(m *testing.M) {
	var err error
	testTaskID, err = Task.Create(&request.Task{
		Name:        "test-" + rand.HexStr(32),
		RunOn:       "",
		BashContent: "sleep 1",
		Description: "",
	})
	if err != nil {
		panic(err)
	}

	m.Run()
}
