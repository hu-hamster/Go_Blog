package router

import (
	"go_blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")
	e.GET("/", controller.Index)
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)
	e.GET("/post_index", controller.GetPostIndex)
	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)
	e.GET("/post_detail", controller.PostDetail)
	e.Run(":8888")
}
