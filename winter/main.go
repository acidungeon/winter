package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct {
	Id     int;
	Name   string;
	Password string;
	FollowId []int;
	SubscriptionId []int;
	coins int;
}
var Users []User;
var State = make(map[string]interface{});//储存状态；

func Register(ctx *gin.Context) {
	name := ctx.Request.FormValue("Name")//前端使用Form；
	password := ctx.Request.FormValue("Password")
	if IsExist(name) {
		State["state"] = 1
		State["text"] = "此用户已存在！"
	} else {
		AddStruct(name, password);
		State["state"] = 1
		State["text"] = "注册成功！"
	}
	ctx.String(http.StatusOK, "%v", State)
}
func Login(ctx *gin.Context) {
	name := ctx.Request.FormValue("Name")
	password := ctx.Request.FormValue("Password")
	if IsExist(name) {
		Password := IsRight(name, password)
		if Password {
			State["state"] = 1
			State["text"] = "登录成功！"
		} else {
			State["state"] = 0
			State["text"] = "密码错误！"
		}
	} else {
		State["state"] = 2
		State["text"] = "登录失败！此用户尚未注册！"
	}

	ctx.String(http.StatusOK, "%v", State)
}
func AddStruct(name string, passwd string) {
	var user User;
	user.Name = name;
	user.Password = passwd;
	user.Id = len(Users) + 1;
	user.FollowId=make([]int,0);
	user.SubscriptionId=make([]int,0);
	user.coins=0;
	Users = append(Users, user);
}//添加用户函数；
func IsExist(user string) bool {
	if len(Users) == 0 {
		return false
	} else {
		for _, i := range Users {
			if i.Name == user {
				return true
			}
		}
	}
	return false
}//判断用户是否存在函数
func IsRight(user string, password string) bool {
	for _, i := range Users {
		if i.Name == user {
			return i.Password == password
		}
	}
	return false
}//判断密码是否正确函数；

func main()  {
	router := gin.Default();
	action := router.Group("admin");
	{
		action.GET("/register", Register);
		action.GET("/login", Login);
	}
	router.Run(":8848")
}