package models

import (
	"gitee.com/favefan/doconsole/global"
	"gorm.io/gorm"
)

type User struct {
	Model
	// Name string `json:`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int    `json:"is_admin"`
	// OwnedContainer string `json:""`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := global.GDB.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}

	return false, nil
}

// GetUserById return user details by given user id
func GetUserById(id int) (*User, error) {
	var user User
	err := global.GDB.Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := global.GDB.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
