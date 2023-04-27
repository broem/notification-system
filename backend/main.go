package main

import (
	"github.com/broem/notification-system/backend/internal/config"
	"github.com/broem/notification-system/backend/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	server.Start(cfg)
}
