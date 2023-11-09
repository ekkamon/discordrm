package databases

import "gorm.io/gorm"

type Conn struct {
	PgSql *gorm.DB
}

func NewMultiDBConnection(pgsql *gorm.DB) *Conn {
	return &Conn{
		PgSql: pgsql,
	}
}
