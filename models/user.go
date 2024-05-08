// models/user.go
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
  ID        uint `gorm:"primary_key"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Email     string `json:"email"`
  Password string `json:"password"`
  DeletedAt *time.Time `gorm:"index"`
  CreatedAt time.Time
  UpdatedAt time.Time
}


type Token struct {

  gorm.Model

    UserID int    `json:"userId"`
    Token  string `json:"token"`
}

