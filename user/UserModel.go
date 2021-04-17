package user

import "time"

type User struct {
	Id              int
	LastChangedTime time.Time
	OpenId          string
	VipExpiredTime  time.Time
	FirstLoginTime  time.Time
	LastLoginTime   time.Time
	NickName        string
}
