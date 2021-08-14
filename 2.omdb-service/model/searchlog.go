package model

import "time"

type SearchLog struct {
	ID         string    `json:"id"`
	Request    string    `json:"request"`
	Response   string    `json:"response"`
	TimeCreate time.Time `json:"time_create"`
}
