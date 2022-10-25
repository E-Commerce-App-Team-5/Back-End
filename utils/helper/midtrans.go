package helper

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func OrderMidtrans(orderId string, price int64) *snap.Response {
	midtrans.ServerKey = "SB-Mid-server-kdk2WKRhPwZDIbE8Uthwc-kI"
	midtrans.ClientKey = "SB-Mid-client-iBy7Ib18FnHwH8VL"
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New("SB-Mid-server-kdk2WKRhPwZDIbE8Uthwc-kI", midtrans.Sandbox)
	// orderId := "ORDER-103"

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: price,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapResp, _ := snap.CreateTransaction(req)
	return snapResp
}

func CheckMidtrans(orderId string) *coreapi.TransactionStatusResponse {
	midtrans.ServerKey = "SB-Mid-server-kdk2WKRhPwZDIbE8Uthwc-kI"
	midtrans.ClientKey = "SB-Mid-client-iBy7Ib18FnHwH8VL"
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New("SB-Mid-server-kdk2WKRhPwZDIbE8Uthwc-kI", midtrans.Sandbox)

	res, _ := c.CheckTransaction(orderId)
	return res
}
