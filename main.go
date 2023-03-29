package main

import (
	"IM/conf"
	"IM/routes"
)

func main() {
	conf.InitConfig()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
