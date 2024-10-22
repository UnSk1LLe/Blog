package bootstrap

import (
	"blogProject/pkg/config"
	"blogProject/pkg/database"
	"blogProject/pkg/html"
	"blogProject/pkg/routing"
	"blogProject/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
