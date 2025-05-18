package models

import "time"

type Tweet struct {
	Id        int
	UserId    int
	User      *User
	Text      string
	CreatedAt time.Time
}

func (t *Tweet) SetUser(user *User) *Tweet {
	if user != nil {
		t.User = user
		t.UserId = user.Id
	}

	return t
}

func (t *Tweet) NeedUser() bool {
	return t.User == nil
}

func (t *Tweet) SetUserIdFromUser() {
	if !t.NeedUser() {
		t.UserId = t.User.Id
	}
}
