package entity

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Type      string    `json:"type"`
	Product   []Product `json:"products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryResponseGet struct {
	ID        uint              `gorm:"primarykey" json:"id"`
	Type      string            `json:"type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Product   []ProductResponse `json:"products"`
}

