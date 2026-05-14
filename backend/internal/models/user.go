package models

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"-"`
	Username string `json:"username"`
	FullName string `db:"full_name" json:"fullName"`
}


