package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/GazpachoGit/microservices/payment/config"
	"github.com/GazpachoGit/microservices/payment/internal/ports"
	"github.com/GazpachoGit/microservices_proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server

	//the structure implements service methods which throw error 'not implemented'
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {

	listen, err := net.Listen("tcp", fmt.Sprintf(":/%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
