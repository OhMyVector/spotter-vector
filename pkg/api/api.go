package api

import (
	"github.com/ohmyvector/spotter-vector/pkg/api/vector"
	"github.com/ohmyvector/spotter-vector/pkg/api/websocket"
	"github.com/ohmyvector/spotter-vector/pkg/core/config"
	"github.com/ohmyvector/spotter-vector/pkg/core/server"
)

func Start(cfg *config.Configuration) error {

	e := server.New()

	// Register static assets
	e.Static("/", cfg.App.StaticPath)

	// Register websocket handler
	e.GET("/ws", websocket.Connection)

	// Register vector core commands
	e.POST("/connect", vector.Connect)
	e.GET("/status", vector.GetConnectionStatus)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
