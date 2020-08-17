package entity

import "time"

type Token struct {
	Value     string     `json:"value"`
	UserID    int64      `json:"user_id"`
	DTCreated time.Time  `json:"dt_created"`
	DTUsed    *time.Time `json:"dt_used"`
	DTExpired time.Time  `json:"dt_expired"`
}
