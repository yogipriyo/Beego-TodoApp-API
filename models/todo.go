package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type Category struct {
	Id 				int 	`orm:"column(id);auto"`
	Name 			string	`orm:"column(name);size(255)"`
	Description 	string	`orm:"column(description);size(255)"`
}

type Status struct {
	Id 				int 	`orm:"column(id);auto"`
	Name 			string	`orm:"column(name);size(255)"`
	Description 	string	`orm:"column(description);size(255)"`
}

type Todo struct {
	Id         		int    		`orm:"column(id);auto"`
	Name       		string 		`orm:"column(name);size(255)"`
	Description		string 		`orm:"column(description);size(255)"`
	Datetime		time.Time 	`orm:"column(datetime);auto_now_add;type(datetime)"`
	Category		int 		`orm:"column(id_category)"`
	Status			int 	 	`orm:"column(id_status)"`
	IdUser			int    		`orm:"column(id_user)"`
}

//AddTodo create new todo
func AddTodo(t *Todo) (*Todo, error) {
	o := orm.NewOrm()
	_, err := o.Insert(t)
	if err != nil {
		log.Print(err)
	}
	return t, err
}

// GetAllCategory
func GetAllTodo() ([]*Todo, error) {
	o := orm.NewOrm()
	var result []*Todo
	num, err := o.QueryTable(Todo{}).RelatedSel().All(&result)
	if err != orm.ErrNoRows && num > 0 {
		return result, nil
	} else {
		return nil, err
	}
}

func GetTodo(id int) (*Todo, error){
	o := orm.NewOrm()
	todo := &Todo{Id: id}

	err := o.Read(todo)
	if err == orm.ErrNoRows {
	    return nil, err
	} else if err == orm.ErrMissPK {
	    return nil, err
	} else {
	    return todo, nil
	}
}

func DeleteTodo(Uid int) (string, error){
	o := orm.NewOrm()
	if _, err := o.Delete(&Todo{Id: Uid}); err == nil {
		return "Todo deleted", err
	} else {
		return "Failed to delete to do", err
	}
}

// Find todo by params
func FindTodo(id int, condition string) ([]*Todo, error) {
	o := orm.NewOrm()
	var todo []*Todo
	_, err := o.QueryTable("todo").Filter(condition, id).All(&todo)
	return todo, err
}

func FindTodoTime(EndTime, StartTime string) ([]*Todo, error){
	o := orm.NewOrm()
	var todo []*Todo
	var err error
	
	if EndTime == ""{
		_, err = o.QueryTable("todo").Filter("datetime__gte", StartTime).All(&todo)
	} else {
		_, err = o.QueryTable("todo").Filter("datetime__gte", StartTime).Filter("datetime__lte", EndTime).All(&todo)
	}

	if err == nil {
		return todo, err
	} else {
		return todo, err
	}
}