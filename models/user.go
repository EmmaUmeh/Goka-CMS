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
  DeletedAt *time.Time `gorm:"index"`
  CreatedAt time.Time
  UpdatedAt time.Time
}


type Token struct {

  gorm.Model

    UserID int    `json:"userId"`
    Token  string `json:"token"`
}

func (user *User) CreateUser(db *gorm.DB) error {
  return db.Create(user).Error
};


func UserExistsByEmail(db *gorm.DB, email string) bool {
  var user User
  // Query the database for a user with the specified email
  if err := db.Where("email = ?", email).First(&user).Error; err != nil {
      // User with the email does not existcx
      return false
  }
  // User with the email exists
  return true
}
