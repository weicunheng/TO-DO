package main

import (
	"TO-DO/controllers"
	_ "TO-DO/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTask;post:NewTask")
	beego.Run()
}

