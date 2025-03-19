package models

import (
	"log/slog"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string `gorm:"size:255"`
	Content string `gorm:"type:text"`
}

func BlogsAll() *[]Blog {
	slog.Info("Fetching all blogs")
	var blogs []Blog
	if err := DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&blogs).Error; err != nil {
		slog.Error("Failed to fetch blogs", slog.String("error", err.Error()))
		return nil
	}
	slog.Info("Successfully fetched blogs", slog.Int("count", len(blogs)))
	return &blogs
}

func BlogsFind(id uint64) *Blog {
	slog.Info("Fetching blog by ID", slog.Uint64("id", id))
	var blog Blog
	if err := DB.Where("id = ?", id).First(&blog).Error; err != nil {
		slog.Error("Failed to fetch blog", slog.Uint64("id", id), slog.String("error", err.Error()))
		return nil
	}
	slog.Info("Successfully fetched blog", slog.Uint64("id", id))
	return &blog
}
