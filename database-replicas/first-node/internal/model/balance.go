package model

import "time"

type Balance struct {
	UserID    uint      `json:"userId" gorm:"primaryKey"`
	Balance   float64   `json:"balance" binding:"required,gte=0"`
	CreatedAt time.Time `json:"createdAt"`
}
