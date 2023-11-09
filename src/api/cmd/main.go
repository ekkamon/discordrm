package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
	"discordrm/api/pkg/databases"
)

func main() {
	// load configs
	cfg := config.LoadConfig()

	// connection database
	pgsql := databases.NewPostgreSQLConnection(cfg)

	// create multiple database
	db := databases.NewMultiDBConnection(pgsql)

	// start fiber
	s := server.NewServer(cfg, db)
	s.Start()
}
