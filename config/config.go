package config

import (
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/parser"
	"github.com/ismdeep/rand"
	"os"
)

// WorkDir working directory ALCHEMY_FURNACE_ROOT
var WorkDir string

// JWT key ALCHEMY_FURNACE_JWT
var JWT string

// Bind listen ALCHEMY_FURNACE_BIND
var Bind string

// EnableSignUp enable sign up ALCHEMY_FURNACE_SIGN_UP_ENABLED
var EnableSignUp bool

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

	// 4. 获取注册启用标记
	EnableSignUp = false
	if os.Getenv("ALCHEMY_FURNACE_SIGN_UP_ENABLED") != "" {
		f, err := parser.ToBool(os.Getenv("ALCHEMY_FURNACE_SIGN_UP_ENABLED"))
		if err == nil && f {
			EnableSignUp = true
		}
	}

	jwt.Init(&jwt.Config{
		Key:    JWT,
		Expire: "72h",
	})
}
