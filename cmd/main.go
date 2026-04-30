package main

import (
	"belajar-go/internal/configs"
	"belajar-go/internal/handler/memberships"
	membershipRepo "belajar-go/internal/repository/memberships"
	membershipSvc "belajar-go/internal/service/memberships"
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
	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
