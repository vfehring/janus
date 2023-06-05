package models

type Guild struct {
	ID              int64 `gorm:"primaryKey"`
	Premium         bool
	AdminRole       int64
	AdminChannel    int64
	LogChannel      int64
	WelcomeChannel  int64
	MaleRole        int64
	FemaleRole      int64
	NonBinaryRole   int64
	GenderfluidRole int64
	MtFRole         int64
	FtMRole         int64
	VerifiedRole    int64
	RequiredRole    int64
	RuleMessage     int64
	RuleRole        int64
}
