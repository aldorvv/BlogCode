package models

import "time"

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`

	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
