package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"testing"
)

var testUserID uint
var testTaskID uint

func TestMain(m *testing.M) {
	var err error
	testUserID, err = User.Register(rand.Username(), rand.PasswordEasyToRemember(4))
	if err != nil {
		panic(err)
	}

	testTaskID, err = Task.Create(testUserID, &request.Task{
		Name:        "test-" + rand.HexStr(32),
		RunOn:       "",
		BashContent: "sleep 1",
		Description: "",
	})

	m.Run()
}
