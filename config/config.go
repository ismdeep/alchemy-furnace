package config

import (
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/rand"
	"os"
)

// WorkDir working directory ALCHEMY_FURNACE_ROOT
var WorkDir string

// JWT key ALCHEMY_FURNACE_JWT
var JWT string

// Bind listen ALCHEMY_FURNACE_BIND
var Bind string

func init() {
	// 1. 获取工作目录
	WorkDir, _ = os.Getwd()
	if os.Getenv("ALCHEMY_FURNACE_ROOT") != "" {
		WorkDir = os.Getenv("ALCHEMY_FURNACE_ROOT")
	}

	// 2. 获取 JWT 密钥
	JWT = rand.Str(32)
	if os.Getenv("ALCHEMY_FURNACE_JWT") != "" {
		JWT = os.Getenv("ALCHEMY_FURNACE_JWT")
	}

	// 3. 获取 Bind 地址
	Bind = "0.0.0.0:8000"
	if os.Getenv("ALCHEMY_FURNACE_BIND") != "" {
		Bind = os.Getenv("ALCHEMY_FURNACE_BIND")
	}

	jwt.Init(&jwt.Config{
		Key:    JWT,
		Expire: "72h",
	})
}
