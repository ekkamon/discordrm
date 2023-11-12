package main

import (
	"discordrm/api/config"
	"discordrm/api/interval/server"
	"discordrm/api/pkg/databases"
	"discordrm/api/pkg/redis"
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

	// redis client
	redis := redis.NewClient(cfg)

	// connection database
	pgsql := databases.NewPostgreSQLConnection(cfg)

	// create multiple database
	db := databases.NewMultiDBConnection(pgsql)

	// start fiber
	s := server.NewServer(cfg, db, redis)
	s.Start()
}
