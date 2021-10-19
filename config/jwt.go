package config

import (
	"leek-api/pkg/config"
)

func init() {
	config.Add("jwt", config.StrMap{
		// 密钥
		"secret": config.Env("jwt.secret", ""),

		// 有效时间，单位：秒
		"ttl": config.Env("jwt.ttl", 3600),
	})
}
