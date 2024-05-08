// models/task.go
package models

import (
	"time"

)

type Task struct {
  ID        uint `gorm:"primary_key"`
  Title string `json:"title"`
  Body  string `json:"body"`
  CreatedAt time.Time
  UpdatedAt time.Time
	User  User   `json:"user"`
}
