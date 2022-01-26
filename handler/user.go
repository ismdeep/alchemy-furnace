package handler

import (
	"encoding/json"
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/digest"
	"github.com/ismdeep/jwt"
)

type userHandler struct {
}

var User = &userHandler{}

func (receiver *userHandler) Register(username string, password string) (uint, error) {
	exists, err := model.UserStore.UserExists(username)
	if err != nil {
		return 0, errors.New("system error")
	}
	if exists {
		return 0, errors.New("user already exists")
	}

	item := &model.User{
		Username: username,
		Digest:   digest.Generate(password),
	}
	model.DB.Create(item)

	return item.ID, nil
}

func (receiver *userHandler) Login(username string, password string) (string, error) {
	user, err := model.UserStore.GetUser(username)
	if err != nil {
		return "", err
	}

	if !digest.Verify(user.Digest, password) {
		return "", errors.New("password verification failed")
	}

	loginUser := schema.LoginUser{
		ID:       user.ID,
		Username: user.Username,
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