package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey"`
	Username  string    `gorm:"type:VARCHAR(128);not null;index:idx_username,type:btree;<-"`
	Password  string    `gorm:"type:VARCHAR(256);not null;<-"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
