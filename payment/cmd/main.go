package main

import (
	"log"

	"github.com/GazpachoGit/microservices/payment/config"
	"github.com/GazpachoGit/microservices/payment/internal/adapters/db"
	"github.com/GazpachoGit/microservices/payment/internal/adapters/grpc"
	"github.com/GazpachoGit/microservices/payment/internal/application/core/api"
)

func main() {
	db, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(db)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	defer grpcAdapter.Stop()
	grpcAdapter.Run()
}
