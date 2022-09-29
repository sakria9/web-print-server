package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sakria9/web-print-server/models"
)

func ServerStatus(c *gin.Context) {
	pending_pages := models.GetPendingPages()
	pending_tasks := models.GetPendingTaskCount()
	c.JSON(200, gin.H{
		"data": gin.H{
			"pending_pages": pending_pages,
			"pending_tasks": pending_tasks,
			"enable_print":  enablePrint,
		},
	})
}
