package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sakria9/web-print-server/models"
)

func ServerStatus(c *gin.Context) {
	pending_pages, err := models.GetPendingPages()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pending_tasks, err := models.GetPendingTaskCount()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": gin.H{
			"pending_pages": pending_pages,
			"pending_tasks": pending_tasks,
			"enable_print":  enablePrint,
		},
	})
}
