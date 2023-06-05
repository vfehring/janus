package models

type User struct {
	ID           string `gorm:"primaryKey"`
	Verified     int
	Premium      bool
	Gender       string
	Rulebreaking int
	Harassment   int
	Catfishing   int
	Underage     int
	Bot          int
}
