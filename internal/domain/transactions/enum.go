package transactions

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)
