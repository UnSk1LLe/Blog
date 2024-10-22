package bootstrap

import (
	"blogProject/internal/database/migration"
	"blogProject/pkg/config"
	"blogProject/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
