package models

type User struct {
	ID string `json:"-"`
	Username string `json:"username"`
	FullName string `db:"full_name" json:"fullName"`
}


