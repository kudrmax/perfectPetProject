package models

type User struct {
	Id           int
	Name         string
	Username     string
	PasswordHash string `json:"-"`
}
