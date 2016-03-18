package routers

import (
	"github.com/yogipriyo/todo_app3/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/todos", &controllers.TodoController{},"get:GetAll;post:Post")
    beego.Router("todos/:id", &controllers.TodoController{},"get:Get;delete:Delete")
    beego.Router("/todos/:condition/:id", &controllers.TodoController{},"get:SearchByCondition")
    beego.Router("/todos/datetime", &controllers.TodoController{},"get:SearchByDatetime")


    beego.Router("/users", &controllers.UserController{},"get:GetAll;post:Post")
    beego.Router("/users/login", &controllers.UserController{},"post:Login")
    beego.Router("/users/logout", &controllers.UserController{},"post:Logout")
    beego.Router("/users/:id",&controllers.UserController{},"get:Get;delete:Delete")
}
