package postgres

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.Installations.EnablePostgreSQLDB = true
	container.InjectInstallation(func() error {
		username := config.DATABASE_POSTGRES_USERNAME.Get()
		pwd := config.DATABASE_POSTGRES_PASSWORD.Get()
		host := config.DATABASE_POSTGRES_HOSTNAME.Get()
		dbname := config.DATABASE_POSTGRES_NAME.Get()
		sslMode := config.DATABASE_POSTGRES_SSL_MODE.Get()
		db, err := gorm.Open(postgres.Open("postgres://" + username + ":" + pwd + "@" + host + "/" + dbname + "?sslmode=" + sslMode))
		if err != nil {
			slog.Logger.Error("error getting postgresql connection", constants.ERROR, err.Error())
			return err
		}
		DB = db
		return nil
	})
}
