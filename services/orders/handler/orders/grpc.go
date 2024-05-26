package handler

import (
	"context"
	"grpc-microservice/services/common/genproto/orders"
	"grpc-microservice/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrder(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}