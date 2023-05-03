package models

type Name struct {
	UID       string `json:"uid" db:"uid"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}
