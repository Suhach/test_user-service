package main

import (
	"github.com/Suhach/test_user-service/internal/database"
	"github.com/Suhach/test_user-service/internal/transport/grpc"
	"github.com/Suhach/test_user-service/internal/user"
)

func main() {
	database.InitDB()

	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)
	handler := grpc.NewHandler(svc)

	grpc.RunServer(handler, ":50051")
}
