package service

import (
	"midtrans-go/model/web"

	"github.com/gin-gonic/gin"
)

type MidtransService interface {
	Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse
}
