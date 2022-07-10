package user

import "time"

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
	Birthday time.Time
	Gender   byte
}

// IsAdult returns whether the user is 18 years or older
func (u User) IsAdult() bool {
	return time.Time(u.Birthday).AddDate(18, 0, 0).Before(time.Now())
}
