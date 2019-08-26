package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type Task struct {
	Id int
	Title string
	Done bool 	`orm:"null";type(bool);default(0)`
	DoneAt time.Time `orm:"null";"auto_now_add;type(datetime)";`
	CreateAt time.Time `orm:"null"`
}

func CreateTask(title string) (taskObj *Task)  {
	o := orm.NewOrm()
	o.Using("default")

	task := new(Task) // {title:title, done:false, create_at: time.Now()}
	//task.ID ID= 1
	task.Title = title
	task.Done = false
	task.CreateAt = time.Now()
	//task.done_at  /= time.Now()
	_ , err := o.Insert(task)
	if err != nil{
		fmt.Println("errrorrrrr", err)
	}
	
	return task
}

func GetAllTask() (tasks Task){
	o := orm.NewOrm()
	task := new(Task)
	err := o.Read(task)
	if err != nil{
		return
	}
	return
}

func init()  {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=todo host=127.0.0.1 port=5432 sslmode=disable", 30)
	orm.RegisterModel(new(Task))
	orm.RunSyncdb("default", false, true)
}