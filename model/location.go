package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (location *Location) BeforeCreate(tx *gorm.DB) (err error) {
	location.ID = uuid.New()
	location.CreatedAt = time.Now()
	location.UpdatedAt = time.Now()
	return
}

func (Location) TableName() string {
	return "locations"
}
