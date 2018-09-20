package models

import "log"

// table users
type User struct {
	BaseModel
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	IsAdmin     string `json:"is_admin"`
	Mark        string `json:"mark"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
}

//模型自动绑定
func (self *User) _listUser() ([]User, error) {
	var users []User
	err := DB.First(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

//查询ID的方式
func (self *User) _GetId(id string) error {

	err := DB.First(&self, id).Error

	if err != nil {
		return err
	}

	return nil

}

func GetAllUser() []User {
	var user User

	users, err := user._listUser()

	if err != nil {
		log.Println(err)
	}

	return users
}

func GetIdUser(id string) User {
	var user User

	err := user._GetId(id)

	if err != nil {
		log.Println(err)
	}

	return user
}
