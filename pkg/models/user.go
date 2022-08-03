package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u User) HashPassword(p string) (hp string, err error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	hp = string(b)

	return
}
