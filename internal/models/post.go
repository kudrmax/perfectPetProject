package models

import "time"

type Post struct {
	Id       int64
	UserId   int64
	Text     string
	Datetime time.Time
}
