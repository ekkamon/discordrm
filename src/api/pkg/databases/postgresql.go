package databases

import (
	"discordrm/api/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLConnection(cfg *config.Config) *gorm.DB {
	pgsqlConnURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.PgSQL.Host,
		cfg.PgSQL.Port,
		cfg.PgSQL.Username,
		cfg.PgSQL.Password,
		cfg.PgSQL.DBName,
		cfg.PgSQL.SSLMode,
		cfg.PgSQL.Timezone,
	)

	db, err := gorm.Open(postgres.Open(pgsqlConnURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("[PostgreSQL] has been connected.")

	return db
}
