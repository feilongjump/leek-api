package bootstrap

import (
	"gorm.io/gorm"
	"leek-api/app/models/article"
	"leek-api/app/models/product"
	"leek-api/app/models/user"
	"leek-api/pkg/config"
	"leek-api/pkg/model"
	"log"
	"time"
)

// SetupDB 数据库初始化
func SetupDB() {
	db := model.ConnectDB()

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	migration(db)
}

func migration(db *gorm.DB) {
	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&article.Content{},
		&product.Product{},
		&product.Sku{},
	)
}
