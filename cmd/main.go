package main

import (
	"sub-aggregator/cfg"
	"sub-aggregator/internal/subaggr"
	"sub-aggregator/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := cfg.LoadConfig()
	database := db.NewDb(conf)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	repo := subaggr.NewRepository(database)
	serv := subaggr.NewService(&subaggr.ServiceDeps{
		Repository: repo,
	})

	subaggr.NewHandler(router, &subaggr.HandlerDeps{
		Config:  conf,
		Service: serv,
	})
}
