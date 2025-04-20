package models

type User struct {
	ID    uint   `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
