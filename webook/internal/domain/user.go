package domain

import "time"

// User领域对象 entity
type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
}
