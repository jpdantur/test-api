package transactions_test

import (
	"errors"
	"testing"

	"github.com/jpdantur/test-api/internal/domain/transactions"
	"github.com/stretchr/testify/assert"
)

func TestService_AddConcurrent(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	res1, err1 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Credit,
		Amount: 1,
	})
	res2, err2 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Debit,
		Amount: 0.5,
	})
	assert.Nil(err1)
	assert.Nil(err2)
	assert.Equal(transactions.Credit, res1.Type)
	assert.Equal(float64(1), res1.Amount)
	assert.Equal(transactions.Debit, res2.Type)
	assert.Equal(0.5, res2.Amount)
}

func TestService_AddInvalidType(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	_, err := service.Add(&transactions.TransactionBody{
		Type:   "foo",
		Amount: -0.5,
	})
	assert.True(errors.Is(err, transactions.ErrInvalidTransactionType))
}

func TestService_AddNegativeBalance(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	_, err := service.Add(&transactions.TransactionBody{
		Type:   transactions.Credit,
		Amount: 1,
	})
	assert.Nil(err)
	_, err = service.Add(&transactions.TransactionBody{
		Type:   transactions.Debit,
		Amount: 1.5,
	})
	assert.True(errors.Is(err, transactions.ErrNegativeBalance))
}

func TestService_Get(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	res, err := service.Add(&transactions.TransactionBody{
		Type:   transactions.Credit,
		Amount: 1,
	})
	assert.Nil(err)
	res2, err := service.GetByID(res.ID)
	assert.Nil(err)
	assert.Equal(res, res2)
}

func TestService_GetInvalidID(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	_, err := service.GetByID("614c0765-dd1f-4fad-8e56-1df1ab8b119")
	assert.True(errors.Is(err, transactions.ErrInvalidID))
}

func TestService_GetNotFound(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	_, err := service.GetByID("614c0765-dd1f-4fad-8e56-1df1ab8b119c")
	assert.True(errors.Is(err, transactions.ErrNotFound))
}

func TestService_GetAll(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	res1, err1 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Credit,
		Amount: 1,
	})
	assert.Nil(err1)
	res2, err2 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Debit,
		Amount: 0.5,
	})
	assert.Nil(err2)
	res := service.GetAll()
	assert.Contains(res, res1)
	assert.Contains(res, res2)
	assert.Len(res, 2)
}

func TestService_GetBalance(t *testing.T) {
	assert := assert.New(t)
	service := transactions.NewService()
	_, err1 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Credit,
		Amount: 1,
	})
	assert.Nil(err1)
	_, err2 := service.Add(&transactions.TransactionBody{
		Type:   transactions.Debit,
		Amount: 0.5,
	})
	assert.Nil(err2)
	res := service.GetBalance()
	assert.Equal(0.5, res)
}
