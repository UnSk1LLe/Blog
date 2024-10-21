package bootstrap

import (
	"blogProject/pkg/config"
	"blogProject/pkg/html"
	"blogProject/pkg/routing"
	"blogProject/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
