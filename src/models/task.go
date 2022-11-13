package models

import "time"

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
}

// json:"xxxxx"はレスポンスのキー