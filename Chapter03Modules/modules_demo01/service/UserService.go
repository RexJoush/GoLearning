package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"modules_demo01/dao"
	"modules_demo01/entities"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	fmt.Println(c)
}

func GetUser(c *gin.Context) {
	param := c.Param("id")
	var users []entities.User
	if param != "" {
		id, err := strconv.Atoi(param)
		if err != nil {
			panic("请输入正确的id值")
		}
		users = append(users, dao.GetUserById(id))
	} else {
		users = dao.GetUser()
	}
	c.AsciiJSON(http.StatusOK, users)
}
func DeleteUser(c *gin.Context) {
	param := c.Param("id")
	if param != "" {
		id, err := strconv.Atoi(param)
		if err != nil {
			panic("请输入正确的id值")
		}
		dao.DeleteUser(id)
	}
	c.AsciiJSON(http.StatusOK, "ok")
}

func EditUser(c *gin.Context) {

}
