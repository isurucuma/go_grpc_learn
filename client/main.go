package main

import (
	"context"
	"github.com/isurucuma/go_grpc_learn/protos/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	client := gen.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := client.GetMenu(ctx, &gen.GetMenuRequest{})
	if err != nil {
		log.Fatalf("could not get menu: %v", err)
	}

	done := make(chan bool)

	var items []*gen.Item
	go func() {
		for {
			menuRes, err := menuStream.Recv()
			if err != nil {
				if err == io.EOF {
					done <- true
					return
				}
				log.Fatalf("could not receive menu: %v", err)
			}

			items = menuRes.Items
			log.Printf("received menu: %v", items)
		}
	}()

	<-done

	receipt, _ := client.PlaceOrder(ctx, &gen.OrderRequest{
		Items: items,
	})

	log.Printf("received receipt: %v", receipt)

	orderStatus, _ := client.GetOrderStatus(ctx, &gen.Receipt{
		Id: receipt.Id,
	})

	log.Printf("received order status: %v", orderStatus)

}
