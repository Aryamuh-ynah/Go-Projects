package models

import "gorm.io/gorm"

type URL struct {

	gorm.Model
	OriginalURL string `gorm:"not null" json:"original_url"`
	ShortCode   string `gorm:"unique;not null" json:"short_code"`

	Clicks int `gorm:"default:0" json:"clicks"`


}
