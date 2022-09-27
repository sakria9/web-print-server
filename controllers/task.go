package controllers

import (
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sakria9/web-print-server/middlewares"
	"github.com/sakria9/web-print-server/models"
	"github.com/sakria9/web-print-server/utils"
)

func CreateTask(c *gin.Context) {
	session := sessions.Default(c)
	var task models.Task
	var user models.User
	task.Email = session.Get(middlewares.UserKey).(string)
	task.Status = models.Pending
	user.Email = task.Email
	if err := user.GetByEmail(); err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(uuid.New().String() + ".pdf")
	task.File = filename
	realPath := utils.GetRealPath(filename)
	if err := c.SaveUploadedFile(file, realPath); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pageNum, err := utils.GetPdfPageNum(realPath)
	if err != nil {
		os.Remove(realPath)
		c.JSON(400, gin.H{"error": "PDF file error"})
		return
	}
	if pageNum > int(user.MaxPage) {
		os.Remove(realPath)
		c.JSON(400, gin.H{"error": "Page number exceed"})
		return
	}
	task.Pages = pageNum

	if err := task.Create(); err != nil {
		os.Remove(realPath)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task created", "id": task.ID})
}

func CancelTask(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get(middlewares.UserKey).(string)
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := task.GetByID(); err != nil {
		c.JSON(400, gin.H{"error": "Task not found"})
		return
	}
	if task.Email != userEmail {
		c.JSON(400, gin.H{"error": "Permission denied"})
		return
	}
	if err := task.TryCancelTask(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": task.Status})
}

func ListTask(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get(middlewares.UserKey).(string)
	tasks, err := models.GetTasksByEmail(userEmail)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}
