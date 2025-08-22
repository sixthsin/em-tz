package main

import (
	"log"
	"sub-aggregator/cfg"
	"sub-aggregator/internal/subs"
	"sub-aggregator/migrations"
	"sub-aggregator/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := cfg.LoadConfig()
	database := db.NewDb(conf)

	migrations.AutoMigrate()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	repo := subs.NewRepository(database)
	serv := subs.NewService(&subs.ServiceDeps{
		Repository: repo,
	})

	subs.NewHandler(router, &subs.HandlerDeps{
		Config:  conf,
		Service: serv,
	})

	if err := router.Run(conf.Server.Port); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
