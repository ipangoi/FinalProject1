package models

import "time"

// Todo represents the model for a todo
type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"ID"`
	NamaTodo    string    `gorm:"not null;type:varchar(191)" json:"namaTodo"`
	Description string    `gorm:"type:varchar(191)" json:"description"`
	Tanggal     string    `json:"tanggal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
