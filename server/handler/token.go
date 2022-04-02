package handler

import (
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
	"github.com/ismdeep/rand"
	"time"
)

type tokenHandler struct {
}

var Token = &tokenHandler{}

func (receiver *tokenHandler) Add(req *request.Token) (uint, string, error) {
	if req == nil || req.Name == "" {
		return 0, "", errors.New("bad request")
	}
	var tokens []model.Token
	if err := model.DB.Where("name=?", req.Name).Find(&tokens).Error; err != nil {
		return 0, "", err
	}

	if len(tokens) >= 1 {
		return 0, "", errors.New("already exists token name")
	}

	token := &model.Token{
		Name:      req.Name,
		Key:       rand.StrWithBase("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_", 32),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := model.DB.Create(token).Error; err != nil {
		return 0, "", err
	}

	return token.ID, token.Key, nil
}

func (receiver *tokenHandler) Update(tokenID uint, req *request.Token) error {
	if req == nil || req.Name == "" {
		return errors.New("bad request")
	}

	var tokens []model.Token
	if err := model.DB.Where("id=?", tokenID).Find(&tokens).Error; err != nil {
		return err
	}
	if len(tokens) <= 0 {
		return errors.New("token not found")
	}

	token := tokens[0]
	token.Name = req.Name
	if err := model.DB.Save(&token).Error; err != nil {
		return err
	}

	return nil
}

func (receiver *tokenHandler) List() ([]response.Token, error) {
	var tokens []model.Token
	if err := model.DB.Find(&tokens).Error; err != nil {
		return nil, err
	}

	var results []response.Token
	for _, token := range tokens {
		results = append(results, response.Token{
			ID:        token.ID,
			Name:      token.Name,
			CreatedAt: token.CreatedAt,
		})
	}

	return results, nil
}

func (receiver *tokenHandler) Delete(tokenID uint) error {
	var tokens []model.Token
	if err := model.DB.Where("id=?", tokenID).Find(&tokens).Error; err != nil {
		return err
	}
	if len(tokens) <= 0 {
		return errors.New("token not found")
	}

	token := tokens[0]
	if err := model.DB.Delete(&token).Error; err != nil {
		return err
	}

	return nil
}
