package transactions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpdantur/test-api/internal/domain/transactions"
)

type Controller struct {
	transactionsService transactions.Service
}

func NewController(transactionsService transactions.Service) *Controller {
	return &Controller{transactionsService: transactionsService}
}

//HandleAdd adds a transaction to history
func (ctrl *Controller) HandleAdd(c *gin.Context) {
	var body transactions.TransactionBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	switch res, err := ctrl.transactionsService.Add(&body); {
	case errors.Is(err, transactions.ErrInvalidTransactionType), errors.Is(err, transactions.ErrNegativeBalance):
		c.JSON(http.StatusBadRequest, err.Error())
	default:
		c.JSON(http.StatusCreated, res)
	}
}

//HandleGetByID returns the transaction with given ID
func (ctrl *Controller) HandleGetByID(c *gin.Context) {
	id := c.Param("id")
	switch res, err := ctrl.transactionsService.GetByID(id); {
	case errors.Is(err, transactions.ErrNotFound):
		c.JSON(http.StatusNotFound, err.Error())
	case errors.Is(err, transactions.ErrInvalidID):
		c.JSON(http.StatusBadRequest, err.Error())
	default:
		c.JSON(http.StatusOK, res)
	}
}

//HandleGetHistory returns transactions history
func (ctrl *Controller) HandleGetHistory(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.transactionsService.GetHistory())
}

//HandleGetBalance returns transactions balance
func (ctrl *Controller) HandleGetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.transactionsService.GetBalance())
}
