package services

import (
	"fmt"
	"log"
	"my-project-be/features/order"
	"my-project-be/middlewares"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	model order.OrderModel
	midtrans midtrans.Client
}

func OrderService(om order.OrderModel, midtrans midtrans.Client) order.OrderService {
	return &service{model: om, midtrans: midtrans}
}

func (s *service) CreateOrder(newOrder order.Order, token *jwt.Token) (order.Order, error) {
	userID, userNama := middlewares.DecodeToken(token)
	newOrder.OrderUniqueID = fmt.Sprintf("order-%d-%d", userID, time.Now().Unix())
	newOrder.UserID = userID
	result, err := s.model.CreateOrder(newOrder, userID)
	if err != nil {
		return order.Order{}, err
	}
	client := s.midtrans
	log.Println("client", client)
	coreGateway := midtrans.CoreGateway{Client: client}
	req := &midtrans.ChargeReq{
		PaymentType: "bank_transfer",
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.Bank(result.PaymentMethod),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: userNama,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  result.OrderUniqueID,
			GrossAmt: int64(result.TotalPrice),	
		},
	}
	resp, err := coreGateway.Charge(req)
	if err != nil {
		return order.Order{}, err
	}
	result.VANumber = resp.VANumbers[0].VANumber
	result.UserID = userID
	result.Status = resp.TransactionStatus
	result.CreatedAt = resp.TransactionTime
	
	return result, nil
}