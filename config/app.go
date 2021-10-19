package config

import (
	"leek-api/pkg/config"
)

func init() {
	config.Add("app", config.StrMap{
		// 应用名称
		"name": config.Env("app.name", "app"),

		// 当前环境，用以区分多环境
		"env": config.Env("app.env", "production"),

		// 是否进入调试模式
		"debug": config.Env("app.debug", false),

		// 应用服务端口
		"port": config.Env("app.port", "3000"),

		// 用以生成链接
		"url": config.Env("app.url", "http://localhost:3000"),
	})
}
