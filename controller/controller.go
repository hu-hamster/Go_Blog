package controller

import (
	"fmt"
	"go_blog/dao"
	"go_blog/model"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := &model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(user)
	c.Redirect(http.StatusMovedPermanently, "/") //跳转到首页
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(http.StatusOK, "login.html", "用户不存在")
	} else {
		if u.Password != password {
			log.Println("密码错误")
			c.HTML(http.StatusOK, "login.html", "密码错误")
		} else {
			log.Println("登录成功")
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(http.StatusOK, "postIndex.html", posts)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	context := c.PostForm("content")
	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: context,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(http.StatusMovedPermanently, "/post_index")
}

//跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))
	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
