package config

import "leek-api/pkg/config"

func init() {

	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{

			// 数据库连接信息
			"host":     config.Env("database.host", "127.0.0.1"),
			"port":     config.Env("database.port", "3306"),
			"database": config.Env("database.database", ""),
			"username": config.Env("database.username", ""),
			"password": config.Env("database.password", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Env("database.max_idle_connections", 100),
			"max_open_connections": config.Env("database.max_open_connections", 25),
			"max_life_seconds":     config.Env("database.max_life_seconds", 5*60),
		},
	})
}
