package handler

import (
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
	"time"
)

type nodeHandler struct{}

var Node = &nodeHandler{}

// Add a node
// @return nodeID uint
// @return err error
func (receiver *nodeHandler) Add(userID uint, req *request.Node) (uint, error) {
	if req == nil {
		return 0, errors.New("req is nil")
	}

	// 1. 检查是否有重名节点
	var cnt int64
	if err := model.DB.Model(&model.Node{}).Where("user_id=? AND name=?", userID, req.Name).Count(&cnt).Error; err != nil {
		return 0, err
	}
	if cnt > 0 {
		return 0, errors.New("already has this node name")
	}

	node := &model.Node{
		UserID:    userID,
		Name:      req.Name,
		Host:      req.Host,
		Port:      req.Port,
		Username:  req.Username,
		SSHKey:    req.SSHKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := model.DB.Create(node).Error; err != nil {
		return 0, errors.New("node info save failed")
	}

	return node.ID, nil
}

// Update a node
func (receiver *nodeHandler) Update(userID uint, nodeID uint, req *request.Node) error {
	if req == nil {
		return errors.New("req is nil")
	}

	// 1. 获取信息
	node := &model.Node{}
	if err := model.DB.Where("id=? AND user_id=?", nodeID, userID).First(node).Error; err != nil {
		return err
	}

	node.Name = req.Name
	node.Host = req.Host
	node.Port = req.Port
	node.Username = req.Username
	node.UpdatedAt = time.Now()
	if req.SSHKey != "" {
		node.SSHKey = req.SSHKey
	}
	if err := model.DB.Save(node).Error; err != nil {
		return err
	}

	return nil
}

// List get nodes
func (receiver *nodeHandler) List(userID uint) ([]response.Node, error) {
	nodes := make([]model.Node, 0)
	if err := model.DB.Where("user_id=?", userID).Find(&nodes).Error; err != nil {
		return nil, err
	}

	var results []response.Node
	for _, node := range nodes {
		results = append(results, response.Node{
			ID:       node.ID,
			Name:     node.Name,
			Host:     node.Host,
			Port:     node.Port,
			Username: node.Username,
		})
	}

	return results, nil
}

func (receiver *nodeHandler) Delete(nodeID uint) error {
	var nodes []model.Node
	if err := model.DB.Where("id=?", nodeID).Find(&nodes).Error; err != nil {
		return err
	}

	if len(nodes) <= 0 {
		return errors.New("node not found")
	}

	node := nodes[0]

	if err := model.DB.Delete(&node).Error; err != nil {
		return err
	}

	return nil
}
