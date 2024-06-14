package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	OriginalURL string `json:"original_url"`
	ShortCode   string `gorm:"uniqueIndex" json:"short_code"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
