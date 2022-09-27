package main

import (
	"github.com/sakria9/web-print-server/db"
	"github.com/sakria9/web-print-server/models"
	"github.com/sakria9/web-print-server/server"
)

func main() {
	db.Init()
	models.Migration()
	server.Init()
}
