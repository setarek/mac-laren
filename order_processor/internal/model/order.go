package model

import "time"

type Order struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
