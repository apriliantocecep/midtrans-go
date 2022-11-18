package controller

import "github.com/gin-gonic/gin"

type MidtransController interface {
	Create(c *gin.Context)
}
