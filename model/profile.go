package model

import (
	"time"
)

// Profile [Comment]
type Profile struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Gender   bool      `json:"gender"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Height   float64   `json:"height"`
	Weight   float64   `json:"weight"`
}
