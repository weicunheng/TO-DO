package controllers

import (
	"TO-DO/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Data["Website"] = "To-Do"
	c.Data["Email"] = "1261360292@qq.com"
	o := orm.NewOrm()
	task := new(models.Task)
	err := o.Read(task)
	if err != nil{
		fmt.Println(err)
	}

	c.Data["tasks"] = task
	fmt.Println(c.Data["tasks"])
}


type TaskController struct {
	beego.Controller
}

func (this *TaskController) ListTask()  {
	tasks := models.GetAllTask()
	this.Data["json"] = tasks
	this.ServeJSON()
}
func (this *TaskController) NewTask()  {
	req := struct{ Title string }{}
	fmt.Println(this.GetString("taskName"))
	fmt.Println(this.Ctx.Input.RequestBody, this.Ctx.Input.GetData("taskName"))
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil{
		fmt.Println("任务名不能为空")
		this.Ctx.Output.Status = 400
		this.Ctx.Output.Body([]byte("任务名不能为空"))
		return
	}

	id := models.CreateTask(req.Title)
	fmt.Println("创建成功：", id)
	this.Ctx.Output.Status = 200
	return



}
