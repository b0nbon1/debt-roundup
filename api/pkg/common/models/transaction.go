package models

import "gorm.io/gorm"

type transactionType string

const (
    DEPOSIT  transactionType = "DEPOSIT"
    TRANSFER transactionType = "TRANSFER"
    REPAYMENT transactionType = "REPAYMENT"
	WITHDRAWAL transactionType = "WITHDRAWAL"
)

type Transaction struct {
    gorm.Model
    TransactionType       transactionType `sql:"type:ENUM('DEPOSIT', 'REPAYMENT', 'TRANSFER', 'WITHDRAWAL')" gorm:"column:transaction_type" json:"transaction_type"`
    Description string `json:"description"`
    UserId      int `json:"user_id"`
    Amount      float64 `json:"amount"`
    Status      bool `json:"sttatus"`
}

