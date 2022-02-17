package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_taskHandler_Create(t *testing.T) {
	userID, err := User.Register(rand.Username(), rand.PasswordEasyToRemember(5))
	assert.NoError(t, err)

	type args struct {
		userID uint
		req    *request.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				userID: userID,
				req: &request.Task{
					Name:        "1",
					BashContent: "sleep 10",
					Description: "noting to describe",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Task.Create(tt.args.userID, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_taskHandler_List(t *testing.T) {
	userID, err := User.Register(rand.Username(), rand.PasswordEasyToRemember(5))
	assert.NoError(t, err)

	taskID, err := Task.Create(userID, &request.Task{
		Name:        rand.Username(),
		BashContent: "sleep 1",
		Description: "",
	})
	assert.NoError(t, err)

	t.Logf("got = %v", taskID)

	type args struct {
		userID uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				userID: userID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Task.List(tt.args.userID)
		})
	}
}
