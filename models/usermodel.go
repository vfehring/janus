package models

type User struct {
	ID      int64 `gorm:"primaryKey"`
	Premium bool
	Gender  string
}
