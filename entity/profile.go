package entity

import (
	"time"
)

// Profile [Comment]
type Profile struct {
	ID       int       `xorm:"'id'"`
	Name     string    `xorm:"'name'"`
	Gender   bool      `xorm:"'gender'"`
	Email    string    `xorm:"'email'"`
	Password string    `xorm:"'password'"`
	Birthday time.Time `xorm:"'birthday'"`
	Height   float64   `xorm:"'height'"`
	Weight   float64   `xorm:"'weight'"`
}
