// models/user.go
package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {

  gorm.Model

    ID        int    `json:"id"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Email     string `json:"email"`
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
      // User with the email does not exist
      return false
  }
  // User with the email exists
  return true
}
