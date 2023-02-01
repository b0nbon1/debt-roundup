package models

import "gorm.io/gorm"

type Debt struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Lender      string `json:"lender"`
}

