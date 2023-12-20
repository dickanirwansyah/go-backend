package users

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	Occupations    string
	PasswordHash   string
	AvatarFileName string
	Roles          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
