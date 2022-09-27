package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sakria9/web-print-server/middlewares"
	"github.com/sakria9/web-print-server/models"
)

func checkAdmin(c *gin.Context) bool {
	var operator models.User
	session := sessions.Default(c)
	operator.Email = session.Get(middlewares.UserKey).(string)
	if err := operator.GetByEmail(); err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return false
	}
	if !operator.Admin {
		c.JSON(400, gin.H{"error": "Permission denied"})
		return false
	}
	return true
}

func SetMaxPage(c *gin.Context) {
	if !checkAdmin(c) {
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	maxPage := user.MaxPage
	if err := user.GetByEmail(); err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	user.MaxPage = maxPage
	if err := user.Update(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Update success"})
}

func EnablePrint(c *gin.Context) {
	if !checkAdmin(c) {
		return
	}
	enablePrint = true
	c.JSON(200, gin.H{"message": "Update success"})
}

func DisablePrint(c *gin.Context) {
	if !checkAdmin(c) {
		return
	}
	enablePrint = false
	c.JSON(200, gin.H{"message": "Update success"})
}
