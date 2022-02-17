package handler

import (
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nodeHandler_Add(t *testing.T) {
	type args struct {
		userID uint
		req    *request.Node
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// OK
			name: "",
			args: args{
				userID: 1,
				req: &request.Node{
					Name:     "test-" + rand.HexStr(32),
					Host:     "127.0.0.1",
					Port:     22,
					Username: "root",
					SSHKey: `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`,
				},
			},
			wantErr: false,
		},
		{
			// req is nil
			name: "",
			args: args{
				userID: 1,
				req:    nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Node.Add(tt.args.userID, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_nodeHandler_Update(t *testing.T) {
	userID, err := User.Register(rand.Username(), rand.PasswordEasyToRemember(4))
	assert.NoError(t, err)

	nodeID, err := Node.Add(userID, &request.Node{
		Name:     "test-" + rand.HexStr(32),
		Host:     "127.0.0.1",
		Port:     22,
		Username: "root",
		SSHKey: `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`,
	})
	assert.NoError(t, err)

	type args struct {
		userID uint
		nodeID uint
		req    *request.Node
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// ok
			name: "",
			args: args{
				userID: userID,
				nodeID: nodeID,
				req: &request.Node{
					Name:     "test-" + rand.HexStr(32),
					Host:     "192.168.0.100",
					Port:     22,
					Username: "root",
					SSHKey: `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Node.Update(tt.args.userID, tt.args.nodeID, &request.Node{
				Name:     "",
				Host:     "",
				Port:     0,
				Username: "",
				SSHKey:   "",
			})
			t.Logf("err = %v", err)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
