// models/task.go
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
  ID        uint `gorm:"primary_key"`
  Title string `json:"title"`
  Body  string `json:"body"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

func(task_data *Task) CreateTask(db *gorm.DB)error {
	return db.Create(task_data).Error
}

func(task_data *Task) ListTaskById(db *gorm.DB)error {
	return db.Find(task_data).Error
}
