package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_triggerHandler_Add(t *testing.T) {
	type args struct {
		taskID uint
		req    *request.Trigger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				taskID: testTaskID,
				req: &request.Trigger{
					Name:        "test-trigger-" + rand.HexStr(32),
					Environment: "VERSION=2.1",
					Cron:        "never",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Trigger.Add(tt.args.taskID, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
