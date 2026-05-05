package main

import (
	"belajar-go/internal/configs"
	"belajar-go/internal/handler/memberships"
	"belajar-go/internal/handler/posts"
	membershipRepo "belajar-go/internal/repository/memberships"
	postsRepo "belajar-go/internal/repository/posts"
	membershipSvc "belajar-go/internal/service/memberships"
	postSvc "belajar-go/internal/service/posts"
	"belajar-go/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postsRepo := postsRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postsService := postSvc.NewService(cfg, postsRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	postsHandler := posts.NewHandler(r, postsService)

	membershipHandler.RegisterRoute()
	postsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
