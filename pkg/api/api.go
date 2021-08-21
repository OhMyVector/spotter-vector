package api

import (
	"github.com/ohmyvector/spotter-vector/pkg/api/websocket"
	"github.com/ohmyvector/spotter-vector/pkg/common/config"
	"github.com/ohmyvector/spotter-vector/pkg/common/server"
)

func Start(cfg *config.Configuration) error {

	e := server.New()
	e.Static("/", cfg.App.StaticPath)
	e.GET("/ws", websocket.Connection)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
