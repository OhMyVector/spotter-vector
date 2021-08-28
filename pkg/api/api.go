package api

import (
	"github.com/OhMyVector/spotter-vector/pkg/api/vector"
	"github.com/OhMyVector/spotter-vector/pkg/api/websocket"
	"github.com/OhMyVector/spotter-vector/pkg/core/config"
	"github.com/OhMyVector/spotter-vector/pkg/core/server"
)

func Start(cfg *config.Configuration) error {

	e := server.New()

	// Register static assets
	e.Static("/", cfg.App.StaticPath)

	// Register websocket handler
	e.GET("/ws", websocket.Connection)

	// Register vector core commands
	e.GET("/connect", vector.Connect)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
