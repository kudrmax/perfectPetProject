package models

import "time"

type Tweet struct {
	Id        int
	UserId    int
	User      *User
	Text      string
	CreatedAt time.Time
}

func (t *Tweet) NeedUser() bool {
	return t.User == nil
}
