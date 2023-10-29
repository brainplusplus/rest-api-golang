package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AdminTable struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Admin struct {
	AdminTable
}

type Admins []Admin

func (admin *Admin) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), 12)
	admin.Password = string(hashedPassword[:])
}

func (admin *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(strings.TrimSpace(password)))
}
