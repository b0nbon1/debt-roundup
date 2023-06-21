package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Firstname       string `json:"firstanme"`
    Lastname string `json:"lastname"`
    Email      string `json:"email"`
	Mobileno      string `json:"mobileno"`
    Assets  []Debt `gorm:"foreignKey:loaner;"`
    Loans []Debt `gorm:"foreignKey:loanee;"`
}

