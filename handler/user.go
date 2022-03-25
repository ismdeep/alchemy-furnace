package handler

import (
	"encoding/json"
	"errors"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/jwt"
)

type userHandler struct {
}

var User = &userHandler{}

func (receiver *userHandler) Login(username string, password string) (string, error) {
	if username != config.ROOT.Auth.Username || password != config.ROOT.Auth.Password {
		return "", errors.New("login failed")
	}

	loginUser := schema.LoginUser{
		Username: username,
	}
	bytes, err := json.Marshal(loginUser)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(string(bytes))
	if err != nil {
		return "", err
	}

	return token, nil
}
