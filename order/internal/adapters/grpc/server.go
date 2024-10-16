package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/GazpachoGit/microservices/order/config"
	"github.com/GazpachoGit/microservices/order/internal/ports"
	"github.com/GazpachoGit/microservices_proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server

	//the structure implements service methods which throw error 'not implemented'
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	defer grpcServer.GracefulStop()

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}

}

func (a Adapter) Stop() {
	a.server.Stop()
}
