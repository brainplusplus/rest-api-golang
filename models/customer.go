package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type CustomerTable struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Customer struct {
	CustomerTable
}

type Customers []Customer

func (customer *Customer) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), 12)
	customer.Password = string(hashedPassword[:])
}

// ComparePassword: Used to compare user stored password and  login  password
func (customer *Customer) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(strings.TrimSpace(password)))
}
