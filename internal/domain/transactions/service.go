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
	GetHistory() []*Transaction
	GetBalance() float64
}

type service struct {
	transactionsMap     map[string]*Transaction
	transactionsHistory []*Transaction
	total               float64
	mu                  sync.Mutex
	transactionChannel  chan Transaction
}

func (s *service) addConcurrent(transaction TransactionBody) {
	s.mu.Lock()
	newTransaction := &Transaction{
		ID:            uuid.New().String(),
		Type:          transaction.Type,
		Amount:        transaction.Amount,
		EffectiveDate: time.Now(),
	}
	s.transactionsMap[newTransaction.ID] = newTransaction
	s.transactionsHistory = append(s.transactionsHistory, newTransaction)
	if transaction.Type == Credit {
		s.total += transaction.Amount
	} else {
		s.total -= transaction.Amount
	}
	s.transactionChannel <- *newTransaction
}

//GetByID returns transaction with given id
func (s *service) GetByID(transactionID string) (*Transaction, error) {
	_, err := uuid.Parse(transactionID)
	if err != nil {
		return nil, ErrInvalidID
	}
	if transaction, ok := s.transactionsMap[transactionID]; ok {
		return transaction, nil
	}
	return nil, ErrNotFound
}

//GetHistory returns transactions in chronological order
func (s *service) GetHistory() []*Transaction {
	return s.transactionsHistory
}

//Add adds a transaction to history and updates balance
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
		total:               0,
		transactionsMap:     make(map[string]*Transaction),
		transactionsHistory: make([]*Transaction, 0),
		transactionChannel:  make(chan Transaction, 20),
	}
}
