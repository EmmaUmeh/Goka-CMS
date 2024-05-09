// models/task.go
package models

import (
	"time"

)

type Events struct {
  ID        uint `gorm:"primary_key"`
  Title string `json:"title"`
  Description  string `json:"description"`
  Date *time.Time `json:"date"`
  CreatedAt time.Time
  UpdatedAt time.Time
// User  User   `json:"user"`
}
