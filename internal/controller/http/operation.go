package http

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (h *Handler) OperationConfirm(c *gin.Context) {
	time.Sleep(time.Second * 3)
	c.JSON(200, gin.H{"msg": "ok"})
}
