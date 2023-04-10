package main

import (
	"os"
	"post/config"
	"post/internal/handler"
	"post/internal/repository"
	"post/pkg/log"
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
		micro.Address("127.0.0.1:8082"),
	)
	
	// Register handler
	if err := pb.RegisterPostHandler(srv.Server(),
		handler.NewPostService(repository.Init(cfg))); err != nil {
		log.Fatal(nil, err, "注册handler失败")
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(nil, err, "服务运行失败")
	}
}
