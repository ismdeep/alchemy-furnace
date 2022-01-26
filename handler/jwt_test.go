package handler

import (
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Jwt_Refresh(t *testing.T) {
	username := rand.Username()
	password := rand.PasswordEasyToRemember(5)
	_, err := User.Register(username, password)
	assert.NoError(t, err)

	token, err := User.Login(username, password)
	assert.NoError(t, err)

	_, err = jwt.VerifyToken(token)
	assert.NoError(t, err)

	Jwt.Refresh()

	_, err = jwt.VerifyToken(token)
	assert.Error(t, err)
}

func Benchmark_Jwt_Refresh(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jwt.Refresh()
	}
}
