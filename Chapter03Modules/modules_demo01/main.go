package main

import (
	"github.com/gin-gonic/gin"
	"modules_demo01/db"
	"modules_demo01/service"
)

func main() {

	// 初始化数据库
	db.InitDB()
	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/user", service.GetUser)
	r.GET("/user/:id", service.GetUser)
	r.POST("/user", service.AddUser)
	r.DELETE("/user", service.DeleteUser)
	r.PUT("/user", service.EditUser)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
