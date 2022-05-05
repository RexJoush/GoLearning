package dao

import (
	"modules_demo01/db"
	"modules_demo01/entities"
)

func GetUser() []entities.User {
	var users []entities.User
	db.DB.Raw("select * from easyuser").Scan(&users)
	return users
}

func GetUserById(id int) entities.User {
	var user entities.User
	db.DB.Raw("select * from easyuser where id = ?", id).Scan(&user)
	return user
}

func DeleteUser(id int) bool {
	db.DB.Raw("delete from easyuser where id = ?", id)
	return true
}

func EditUser(u entities.User) bool {
	return true
}
