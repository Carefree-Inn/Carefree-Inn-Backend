package main

import (
	"os"
	"user-post/config"
	"user-post/internal/handler"
	"user-post/internal/repository"
	"user-post/pkg/log"
	pb "user-post/proto"
	
	"go-micro.dev/v4"
)

func main() {
	// Create service
	log.NewLogger("./log/")
	
	cfg := config.Run(os.Getenv("POST_CONFIG_FILE_PATH"))
	
	srv := micro.NewService()
	srv.Init(
		micro.Name(cfg.Micro.Service),
		micro.Version(cfg.Micro.Version),
		micro.Address(cfg.Server.Http.Address),
	)
	
	// Register handler
	if err := pb.RegisterUserPostHandler(srv.Server(),
		handler.NewUserPostService(
			repository.Init(cfg.Database.Dsn))); err != nil {
		log.Fatal(nil, err, "注册handler失败")
	}
	
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(nil, err, "服务运行失败")
	}
}
