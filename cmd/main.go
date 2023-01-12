package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/client"
	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/config"
	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/db"
	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/pb"
	"github.com/Fachrulmustofa20/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed load config: ", err)
	}

	h := db.Init(cfg.DBUrl)

	listen, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	productSvc := client.InitProductServiceClient(cfg.ProductSvcUrl)

	fmt.Println("Order svc on", cfg.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
