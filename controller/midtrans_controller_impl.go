package controller

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransControllerImpl struct {
	MidtransService service.MidtransService
}

func NewMidtransControllerImpl(midtransService service.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *gin.Context) {
	var request web.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.PanicIfError(err)
	}

	midtransResponse := controller.MidtransService.Create(c, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
