package models

import (
	"time"
)

type Account struct {
	ID          uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	Password    string    `json:"password"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
