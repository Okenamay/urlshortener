package models

import "gorm.io/gorm"

const (
	IDSize = 8
)

type URL struct {
	gorm.Model
	URL     string `gorm:"index:idx_url,unique;size:255"`
	ShortID string `gorm:"index:idx_short,unique;size:8"`
}
