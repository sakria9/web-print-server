package main

import (
	"github.com/sakria9/web-print-server/config"
	"github.com/sakria9/web-print-server/db"
	"github.com/sakria9/web-print-server/models"
	"github.com/sakria9/web-print-server/server"
)

func main() {
	config.Init()
	conf := config.GetConfig()
	db.Init(conf.GetString("db"))
	models.Migration()
	server.Init()
}
