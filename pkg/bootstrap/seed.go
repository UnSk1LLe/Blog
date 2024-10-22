package bootstrap

import (
	"blogProject/internal/database/seeder"
	"blogProject/pkg/config"
	"blogProject/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed(database.Connection())
}
