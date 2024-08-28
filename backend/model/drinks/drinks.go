package drinks

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ID        uint           `          gorm:"primarykey" json:"id"`
}

type Drinks struct {
	GormModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
}
