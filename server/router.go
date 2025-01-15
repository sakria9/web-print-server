package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sakria9/web-print-server/controllers"
	"github.com/sakria9/web-print-server/middlewares"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthCookieStore())
	r.POST("/api/logout", controllers.Logout)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/me", middlewares.AuthMiddleware(), controllers.Me)
	r.POST("/api/change-password", middlewares.AuthMiddleware(), controllers.ChangePassword)
	r.GET("/api/server", middlewares.AuthMiddleware(), controllers.ServerStatus)
	task := r.Group("/api/task")
	task.Use(middlewares.AuthMiddleware())
	{
		task.POST("/create", controllers.CreateTask)
		task.POST("/cancel", controllers.CancelTask)
		task.GET("/list", controllers.ListTaskByEmail)
	}
	admin := r.Group("/api/admin")
	admin.Use(middlewares.AuthMiddleware())
	{
		admin.POST("/register", controllers.Register)
		admin.POST("/disable-print", controllers.DisablePrint)
		admin.POST("/enable-print", controllers.EnablePrint)
		admin.POST("/set-max-page", controllers.SetMaxPage)
		admin.GET("/list-task", controllers.ListAllTasks)
	}
	return r
}
