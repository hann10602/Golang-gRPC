package types

import (
	"context"
	"grpc-microservice/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrder(context.Context) []*orders.Order
}