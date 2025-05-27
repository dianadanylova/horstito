package models

import "time"

type Session struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Tag       string    `json:"tag"`
	Notes     string    `json:"notes"`
}
