package config

import (
	"leek-api/pkg/config"
)

func init() {
	config.Add("jwt", config.StrMap{
		// 密钥
		"secret": config.Env("JWT_SECRET"),

		// 有效时间，单位：秒
		"ttl": config.Env("JWT_TTL", 3600),
	})
}
