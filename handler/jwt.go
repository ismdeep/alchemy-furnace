package handler

import (
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/rand"
)

type jwtHandler struct {
}

var Jwt = &jwtHandler{}

func (receiver *jwtHandler) Refresh() {
	jwt.Init(&jwt.Config{
		Key:    rand.HexStr(128),
		Expire: "72h",
	})
}
