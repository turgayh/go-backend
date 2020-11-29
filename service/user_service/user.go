package user_service

import "github.com/EDDYCJY/go-gin-example/models"

type User struct {
	Email    string
	Password string
	Group    int
}

func (a *User) Check() (bool, error) {
	return models.CheckUser(a.Email, a.Password)
}
