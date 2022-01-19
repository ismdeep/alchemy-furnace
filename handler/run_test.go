package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runHandler_List(t *testing.T) {
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				page: 1,
				size: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := Run.List("", tt.args.page, tt.args.size)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
