package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_taskHandler_Create(t *testing.T) {
	type args struct {
		req *request.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				req: &request.Task{
					Name:        "1",
					BashContent: "sleep 10",
					Description: "nothing to describe",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Task.Create(tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_taskHandler_List(t *testing.T) {
	taskID, err := Task.Create(&request.Task{
		Name:        rand.Username(),
		BashContent: "sleep 1",
		Description: "",
	})
	assert.NoError(t, err)

	t.Logf("got = %v", taskID)

	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Task.List()
		})
	}
}
