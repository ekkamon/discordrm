package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	s := server.NewServer(cfg)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
