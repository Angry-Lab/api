package entity

import "time"

type User struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Login        string     `json:"login"`
	Password     string     `json:"password,omitempty"`
	Status       int8       `json:"status"`
	Role         string     `json:"role"`
	DTCreated    time.Time  `json:"dt_created"`
	DTUpdated    time.Time  `json:"dt_updated"`
	DTLastLogged *time.Time `json:"dt_last_logged"`
}
