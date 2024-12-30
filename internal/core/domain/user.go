package domain

import "time"

type User struct {
	ID        string
	Username  string
	Email     string
	ProfileID string
	CreatedAt time.Time
}

type Profile struct {
	ID          string
	UserID      string
	DisplayName string
	Bio         string
	CreatedAt   time.Time
}
