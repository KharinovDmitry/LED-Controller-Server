package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}
