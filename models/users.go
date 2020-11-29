package models

import "github.com/jinzhu/gorm"

type Users struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Group    int    `json:"group"`
}

// CheckAuth checks if authentication information exists
func CheckUser(email, password string) (bool, error) {
	var users Users
	err := db.Select("id").Where(Users{Email: email, Password: password}).First(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if users.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUserPermissionGroup(email string) (int, error) {
	var users Users
	err := db.Select("id").Where(Users{Email: email}).First(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}

	if users.ID > 0 {
		return users.Group, nil
	}

	return -1, nil
}
