package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) HandlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}