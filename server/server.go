package main

import (
	"context"
	"github.com/isurucuma/go_grpc_learn/protos/gen"
)

type server struct {
	gen.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(request *gen.GetMenuRequest, stream gen.CoffeeShop_GetMenuServer) error {
	items := []*gen.Item{
		{
			Id:    "1",
			Name:  "Espresso",
			Price: 123.56,
		},
		{
			Id:    "2",
			Name:  "Latte",
			Price: 234.56,
		},
		{
			Id:    "3",
			Name:  "Cappuccino",
			Price: 345.67,
		},
	}

	for i, _ := range items {
		err := stream.Send(&gen.Menu{
			Items: items[0 : i+1],
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) PlaceOrder(ctx context.Context, request *gen.OrderRequest) (*gen.Receipt, error) {
	return &gen.Receipt{Id: "123"}, nil
}

func (s *server) GetOrderStatus(ctx context.Context, request *gen.Receipt) (*gen.OrderStatus, error) {
	return &gen.OrderStatus{Status: "PREPARING"}, nil
}
