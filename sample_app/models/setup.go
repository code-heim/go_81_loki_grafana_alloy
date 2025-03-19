package models

import (
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	slog.Info("Connecting to database")
	database, err := gorm.Open(mysql.Open("net_http_blogs:tmp_pwd@tcp(127.0.0.1:3306)/net_http_blogs?charset=utf8&parseTime=true"), &gorm.Config{})

	if err != nil {
		slog.Error("Failed to connect to database", slog.String("error", err.Error()))
		panic("Failed to connect to database!")
	}

	DB = database
	slog.Info("Database connection established")
}

func DBMigrate() {
	slog.Info("Starting database migration")
	if err := DB.AutoMigrate(&Blog{}); err != nil {
		slog.Error("Database migration failed", slog.String("error", err.Error()))
		return
	}
	slog.Info("Database migration completed successfully")
}
