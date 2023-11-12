package databases

import (
	"discordrm/api/config"
	"discordrm/api/interval/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(postgres.Open(pgsqlConnURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "PostgreSQL",
		}).Fatal(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "PostgreSQL",
		}).Fatal(err)
	}

	logrus.Info("PostgreSQL has been connected.")

	return db
}
