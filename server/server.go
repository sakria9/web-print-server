package server

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sakria9/web-print-server/config"
	"github.com/sakria9/web-print-server/controllers"
)

func Init() {
	config.Init()
	conf := config.GetConfig()
	r := setupRouter()
	go func() {
		gocron.Every(1).Seconds().Do(controllers.CheckTask)
		<-gocron.Start()
	}()
	r.Run(conf.GetString("port"))
}
