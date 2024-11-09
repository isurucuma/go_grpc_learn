package main

import (
	"context"
	protos2 "github.com/isurucuma/go_grpc_coffeeshop/gen"
)

type server struct {
	protos2.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(request *protos2.GetMenuRequest, stream protos2.CoffeeShop_GetMenuServer) error {
	items := []*protos2.Item{
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
		err := stream.Send(&protos2.Menu{
			Items: items[0 : i+1],
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) PlaceOrder(ctx context.Context, request *protos2.OrderRequest) (*protos2.Receipt, error) {
	return &protos2.Receipt{Id: "123"}, nil
}

func (s *server) GetOrderStatus(ctx context.Context, request *protos2.Receipt) (*protos2.OrderStatus, error) {
	return &protos2.OrderStatus{Status: "PREPARING"}, nil
}
