package routing

import (
	"blogProject/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	r := GetRouter()

	conf := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port))
	if err != nil {
		log.Fatalf("error in routing: %s", err.Error())
		return
	}
}
