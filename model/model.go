package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Year      string         `json:"year"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" sql:"index:deleted_at"`
}
