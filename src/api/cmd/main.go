package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
	"discordrm/api/pkg/databases"
	"discordrm/api/pkg/utils"
)

func main() {
	// load configs
	cfg := config.LoadConfig()

	// new validator instance
	utils.NewValidator()

	// new jwt services
	utils.NewService(cfg)

	// connection database
	pgsql := databases.NewPostgreSQLConnection(cfg)

	// create multiple database
	db := databases.NewMultiDBConnection(pgsql)

	// start fiber
	s := server.NewServer(cfg, db)
	s.Start()
}
