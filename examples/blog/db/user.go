package db

import "fmt"

type UserModel struct {
	Name     string
	Password string
	Address  string
	Gender   int
}

type UserRepository struct {
	Users []*UserModel
}

func (user *UserRepository) GetUserByName(name string) *UserModel {
	for _, model := range user.Users {
		if model.Name == name {
			return model
		}
	}
	return nil
}

func (user *UserRepository) AddUser(model *UserModel) {
	user.Users = append(user.Users, model)
	fmt.Println(model.Name, model.Password, model.Address)
}
