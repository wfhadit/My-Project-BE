package services

import (
	"fmt"
	"log"
	"my-project-be/features/cart"
	"my-project-be/features/order"
	"my-project-be/middlewares"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	model order.OrderModel
	midtrans midtrans.Client
	cart cart.CartModel
}

func OrderService(om order.OrderModel, midtrans midtrans.Client, cart cart.CartModel) order.OrderService {
	return &service{model: om, midtrans: midtrans, cart: cart}
}

func (s *service) CreateOrder(newOrder order.Order, token *jwt.Token) (order.Order, error) {
	userID, userNama := middlewares.DecodeToken(token)
	newOrder.OrderUniqueID = fmt.Sprintf("order-%d-%d", userID, time.Now().Unix())
	newOrder.UserID = userID
	client := s.midtrans
	coreGateway := midtrans.CoreGateway{Client: client}
	req := &midtrans.ChargeReq{
		PaymentType: "bank_transfer",
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.Bank(newOrder.PaymentMethod),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: userNama,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  newOrder.OrderUniqueID,
			GrossAmt: int64(newOrder.TotalPrice),	
		},
	}
	resp, err := coreGateway.Charge(req)
	if err != nil {
		return order.Order{}, err
	}
	newOrder.VANumber = resp.VANumbers[0].VANumber
	newOrder.UserID = userID
	newOrder.Status = resp.TransactionStatus
	newOrder.CreatedAt = resp.TransactionTime
	errDelete := s.cart.DeleteCart(userID)
	if errDelete != nil {
		return order.Order{}, errDelete
	}
	result, err := s.model.CreateOrder(newOrder)
	if err != nil {
		return order.Order{}, err
	}

	
	return result, nil
}

func (s *service) GetOrderByUniqueID(token *jwt.Token) (order.Order, error) {
	userID, _ := middlewares.DecodeToken(token)
	last,err := s.model.GetLastOrder(userID)
	if err != nil {
		return order.Order{}, err
	}
	client := s.midtrans
	coreGateway := midtrans.CoreGateway{Client: client}
	resp, err := coreGateway.Status(last.OrderUniqueID)
	if err != nil {
		return order.Order{}, err
	}
	newStatus:= resp.TransactionStatus
	result, err := s.model.GetOrderByUniqueID(last.OrderUniqueID, userID, newStatus)
	if err != nil {
		return order.Order{}, err
	}
	return result, nil
}

func (s *service) GetAllOrders(token *jwt.Token) ([]order.Order, error) {
	userID, _ := middlewares.DecodeToken(token)
	orders, err := s.model.GetAllOrders(userID)
	if err != nil {
		return []order.Order{}, err
	}
	client := s.midtrans
	coreGateway := midtrans.CoreGateway{Client: client}
	updatedOrders := []order.Order{}
	for _, order := range orders {
		if order.Status == "pending" {
			resp, err := coreGateway.Status(order.OrderUniqueID)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if order.Status != resp.TransactionStatus {order.Status = resp.TransactionStatus
				_, errUpdate := s.model.GetOrderByUniqueID(order.OrderUniqueID, userID, order.Status)
				if errUpdate != nil {
					log.Println(errUpdate.Error())
				}
			}
		}
		updatedOrders = append(updatedOrders, order)
	
	}
	return updatedOrders, nil
}