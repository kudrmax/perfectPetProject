package models

import "time"

type Tweet struct {
	Id        int
	UserId    int
	Text      string
	CreatedAt time.Time
}
