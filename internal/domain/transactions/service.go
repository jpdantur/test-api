package transactions

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound               = errors.New("transaction not found")
	ErrInvalidID              = errors.New("invalid id")
	ErrNegativeBalance        = errors.New("total balance cannot be negative")
	ErrInvalidTransactionType = errors.New("invalid transaction type")
)

type Service interface {
	Add(transaction *TransactionBody) (*Transaction, error)
	GetByID(transactionID string) (*Transaction, error)
	GetAll() []*Transaction
	GetBalance() float64
}

type service struct {
	transactions       map[string]*Transaction
	total              float64
	mu                 sync.Mutex
	transactionChannel chan Transaction
}

func (s *service) addConcurrent(transaction TransactionBody) {
	s.mu.Lock()
	newTransaction := &Transaction{
		ID:            uuid.New().String(),
		Type:          transaction.Type,
		Amount:        transaction.Amount,
		EffectiveDate: time.Now(),
	}
	s.transactions[newTransaction.ID] = newTransaction
	if transaction.Type == Credit {
		s.total += transaction.Amount
	} else {
		s.total -= transaction.Amount
	}
	s.transactionChannel <- *newTransaction
}

func (s *service) GetByID(transactionID string) (*Transaction, error) {
	_, err := uuid.Parse(transactionID)
	if err != nil {
		return nil, ErrInvalidID
	}
	if transaction, ok := s.transactions[transactionID]; ok {
		return transaction, nil
	}
	return nil, ErrNotFound
}

func (s *service) GetAll() []*Transaction {
	transactions := make([]*Transaction, 0, len(s.transactions))
	for _, v := range s.transactions {
		transactions = append(transactions, v)
	}
	return transactions
}

func (s *service) Add(transaction *TransactionBody) (*Transaction, error) {
	if transaction.Type != Debit && transaction.Type != Credit {
		return nil, ErrInvalidTransactionType
	}
	if transaction.Type == Debit && s.total-transaction.Amount < 0 {
		return nil, ErrNegativeBalance
	}
	defer s.mu.Unlock()
	go s.addConcurrent(*transaction)
	response := <-s.transactionChannel
	return &response, nil
}

func (s *service) GetBalance() float64 {
	return s.total
}

func NewService() Service {
	return &service{
		total:              0,
		transactions:       make(map[string]*Transaction),
		transactionChannel: make(chan Transaction, 20),
	}
}
