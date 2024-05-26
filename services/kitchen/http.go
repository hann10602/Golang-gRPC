package main

import (
	"context"
	"grpc-microservice/services/common/genproto/orders"
	"html/template"
	"log"
	"net/http"
	"time"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID: 3123,
			Quantity: 2,
		})

		if err != nil {
			log.Fatalf("Client error1: %v", err)
		}

		res, err := c.GetOrder(ctx, &orders.GetOrderRequest{
			CustomerID: 42,
		})

		if err != nil {
			log.Fatalf("Client error2: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("Template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Kitchen Orders</title>
  </head>
  <body>
    <h1>Kitchen Orders</h1>
    <table border="1">
      <tr>
        <th>Order ID</th>
        <th>Customer ID</th>
        <th>Quantity</th>
      </tr>
      {{range .}}
      <tr>
        <td>{{.OrderID}}</td>
        <td>{{.CustomerID}}</td>
        <td>{{.Quantity}}</td>
      </tr>
	  {{end}}
    </table>
  </body>
</html>
`