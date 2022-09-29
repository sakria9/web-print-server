package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sakria9/web-print-server/middlewares"
	"github.com/sakria9/web-print-server/models"
	"github.com/sakria9/web-print-server/utils"
)

func Login(c *gin.Context) {
	var user models.User
	session := sessions.Default(c)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	password := user.Password
	if err := user.GetByEmail(); err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		c.JSON(400, gin.H{"error": "Wrong password"})
		return
	}

	session.Set(middlewares.UserKey, user.Email)
	session.Save()
	c.JSON(200, gin.H{"message": "Login success", "data": user})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	c.JSON(200, gin.H{"message": "Logout success"})
}

func Register(c *gin.Context) {
	if !checkAdmin(c) {
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Password = hash
	if err := user.Create(); err != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}
	c.JSON(200, gin.H{"message": "Register success", "data": user})
}

func Me(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	user.Email = session.Get(middlewares.UserKey).(string)
	if err := user.GetByEmail(); err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": user})
}
