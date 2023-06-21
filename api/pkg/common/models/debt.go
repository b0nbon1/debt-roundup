package models

import "gorm.io/gorm"

type Debt struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Loaner      int `json:"loaner"`
    Loanee      int `json:"loanee"`
    Amount      float64 `json:"amount"`
    Active      bool `json:"active"`
}

