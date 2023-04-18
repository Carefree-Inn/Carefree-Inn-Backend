package main

import (
	micro "go-micro.dev/v4"
	"os"
	"user/config"
	"user/internal/handler"
	"user/internal/repository"
	log "user/pkg/log"
	pb "user/proto"
)

func main() {
	// Create service
	log.NewLogger("./log/")
	cfg := config.Run(os.Getenv("USER_CONFIG_FILE_PATH"))
	
	srv := micro.NewService()
	srv.Init(
		micro.Name(cfg.Micro.Service),
		micro.Version(cfg.Micro.Version),
		micro.Address(cfg.Server.Http.Address),
	)
	
	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(),
		handler.NewUserService(repository.Init(cfg.Database.Dsn)));
		err != nil {
		log.Fatal(nil, err, "注册handler失败")
	}
	
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(nil, err, "服务运行失败")
	}
}
