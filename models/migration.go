package models

import (
	"github.com/sakria9/web-print-server/db"
)

func Migration() {
	db.GetDB().AutoMigrate(&User{})
	db.GetDB().AutoMigrate(&Task{})
}
