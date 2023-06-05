package models

type Guild struct {
	ID              string `gorm:"primaryKey"`
	Premium         bool
	AdminRole       string
	AdminChannel    string
	LogChannel      string
	WelcomeChannel  string
	MaleRole        string
	FemaleRole      string
	NonBinaryRole   string
	GenderfluidRole string
	MtFRole         string
	FtMRole         string
	VerifiedRole    string
	RequiredRole    string
	RuleMessage     string
	RuleRole        string
}
