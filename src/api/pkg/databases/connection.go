package databases

import (
	"gorm.io/gorm"
)

type Conn struct {
	PgSql *gorm.DB
	Redis *Redis
}

func NewMultiDBConnection(pgsql *gorm.DB, redis *Redis) *Conn {
	return &Conn{
		PgSql: pgsql,
		Redis: redis,
	}
}
