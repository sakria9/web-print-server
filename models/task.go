package models

import (
	"sync"

	"github.com/sakria9/web-print-server/db"
)

type Task struct {
	ID     int64  `json:"id" gorm:"primary_key;autoIncrement"`
	File   string `json:"file"`
	Email  string `json:"email"`
	Pages  int    `json:"pages"`
	Status string `json:"status"`
}

const Pending = "pending"
const Printing = "printing"
const Printed = "printed"
const Cancelled = "cancelled"
const Error = "error"

var lock sync.Mutex

func (t *Task) Create() error {
	return db.GetDB().Create(t).Error
}

func (t *Task) Update() error {
	return db.GetDB().Save(t).Error
}

func (t *Task) GetByID() error {
	return db.GetDB().Where("id = ?", t.ID).First(t).Error
}

func GetPendingTaskCount() (int64, error) {
	var count int64
	err := db.GetDB().Model(&Task{}).Where("status = ?", Pending).Count(&count).Error
	return count, err
}

func GetFirstPendingTask() (*Task, error) {
	lock.Lock()
	defer lock.Unlock()

	cnt, err := GetPendingTaskCount()
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	var task Task
	err = db.GetDB().Where("status = ?", Pending).First(&task).Error
	if err != nil {
		return nil, err
	}
	task.Status = Printing
	err = task.Update()
	if err != nil {
		return nil, err
	}
	return &task, err
}

func (t *Task) TryCancelTask() error {
	lock.Lock()
	defer lock.Unlock()
	err := db.GetDB().Where("id = ?", t.ID).First(t).Error
	if err != nil {
		return err
	}
	if t.Status == Pending {
		t.Status = Cancelled
		return t.Update()
	}
	return nil
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.GetDB().Find(&tasks).Error
	return tasks, err
}

func GetTasksByEmail(email string) ([]Task, error) {
	var tasks []Task
	err := db.GetDB().Where("email = ?", email).Find(&tasks).Error
	return tasks, err
}

func GetTaskCount() (int64, error) {
	var count int64
	err := db.GetDB().Model(&Task{}).Count(&count).Error
	return count, err
}

func (t *Task) GetByEmail() error {
	return db.GetDB().Where("email = ?", t.Email).First(t).Error
}
