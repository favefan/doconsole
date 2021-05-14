package models

import (
	"time"
)

// Model is the default model strcut.
type Model struct {
	Id        uint           `gorm:"primaryKey";form:"Id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

