package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
	"discordrm/api/pkg/databases"
	"discordrm/api/pkg/utils"
)

func main() {
	// setup logger
	utils.NewLogger()

	// load configs
	cfg := config.LoadConfig()

	// new util functions
	utils.NewValidator()
	utils.NewJWT(cfg)

	// connection database
	pgsql := databases.NewPostgreSQLConnection(cfg)
	redis := databases.NewRedisClient(cfg)

	// create multiple database
	db := databases.NewMultiDBConnection(pgsql, redis)

	// start fiber
	s := server.NewServer(cfg, db)
	s.Start()
}
