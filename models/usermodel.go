package models

type User struct {
	ID           int64 `gorm:"primaryKey"`
	Premium      bool
	Gender       string
	Rulebreaking int
	Harassment   int
	Catfishing   int
	Underage     int
	Bot          int
}
