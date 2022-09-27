package controllers

import (
	"github.com/sakria9/web-print-server/models"
	"github.com/sakria9/web-print-server/printer"
	"github.com/sakria9/web-print-server/utils"
)

var currentTask *models.Task

var enablePrint = true

func CheckTask() {
	if !enablePrint {
		return
	}
	empty, err := printer.IsPrintQueueEmpty()
	if err != nil {
		return
	}
	if empty {
		if currentTask != nil {
			currentTask.Status = models.Printed
			currentTask.Update()
			currentTask = nil
		}

		task, err := models.GetFirstPendingTask()
		if err != nil {
			return
		}
		printer.AddToPrintQueue(utils.GetRealPath(task.File))
	}
}
