package handler

import (
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_User_Register(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				username: rand.Username(),
				password: rand.PasswordEasyToRemember(5),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := User.Register(tt.args.username, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Benchmark_User_Register(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := User.Register(rand.Username(), rand.HexStr(32))
		assert.NoError(b, err)
	}
}

func Test_User_Login(t *testing.T) {
	username := rand.Username()
	password := rand.PasswordEasyToRemember(5)
	_, err := User.Register(username, password)
	assert.NoError(t, err)

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				username: username,
				password: password,
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				username: username,
				password: password + rand.HexStr(32),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := User.Login(tt.args.username, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			t.Logf("got = [%v] %v", len(got), got)
		})
	}
}

func Benchmark_User_Login(b *testing.B) {
	username := rand.Username()
	password := rand.PasswordEasyToRemember(5)
	_, err := User.Register(username, password)
	assert.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := User.Login(username, password)
		assert.NoError(b, err)
	}
}
