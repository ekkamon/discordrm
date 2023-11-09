package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
)

func main() {
	// load configs
	cfg := config.LoadConfig()

	// start fiber
	s := server.NewServer(cfg)
	s.Start()
}
