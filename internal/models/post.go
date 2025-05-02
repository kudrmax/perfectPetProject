package models

import "time"

type Post struct {
	Id       int
	UserId   int
	Text     string
	Datetime time.Time
}
