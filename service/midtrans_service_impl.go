package service

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransServiceImpl struct {
	Validate *validator.Validate
}

func NewMidtransServiceImpl(validate *validator.Validate) *MidtransServiceImpl {
	return &MidtransServiceImpl{
		Validate: validate,
	}
}

func (service *MidtransServiceImpl) Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper.PanicIfError(err)
	}

	// request midtrans
	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// user id
	user_id := strconv.Itoa(request.UserId)

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + user_id + "-" + request.ItemID,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		helper.PanicIfError(errSnap.GetRawError())
	}

	midtransReponse := web.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}

	return midtransReponse
}
