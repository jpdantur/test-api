package transactions

import (
	"time"
)

type Transaction struct {
	ID            string          `json:"id"`
	Type          TransactionType `json:"type"`
	Amount        float64         `json:"amount"`
	EffectiveDate time.Time       `json:"effectiveDate"`
}

type TransactionBody struct {
	Type   TransactionType `json:"type" binding:"required"`
	Amount float64         `json:"amount" binding:"required,numeric,min=0"`
}
