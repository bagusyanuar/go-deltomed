package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Division struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (division *Division) BeforeCreate(tx *gorm.DB) (err error) {
	division.ID = uuid.New()
	division.CreatedAt = time.Now()
	division.UpdatedAt = time.Now()
	return
}

func (Division) TableName() string {
	return "divisions"
}
