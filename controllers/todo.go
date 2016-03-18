package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/yogipriyo/todo_app3/models"
	// "time"
	//"log"
	//"fmt"
)

// oprations for Status
type TodoController struct {
	beego.Controller
}

func (t *TodoController) Post() {
	Token := t.GetString("Token")
	if len(Token) > 0{
		Id, Level, Status := models.IsSignedIn(t.GetString("Token"))
	 	if Status {
	 		if Level == 3{
		 		Category, _ := t.GetInt("Category")
				Status, _ := t.GetInt("Status")
				v := models.Todo{Name: t.GetString("Name"), Description: t.GetString("Description"), Category: Category, Status: Status, IdUser: Id}
				todo, err := models.AddTodo(&v)
				if err == nil {
					t.Data["json"] = map[string]interface{}{"message": "Todo successfully saved.", "result": todo}
				} else {
					t.Data["json"] = map[string]interface{}{"message": "Todo unsuccessfully saved.", "error":err}
				}
			} else {
				t.Data["json"] = map[string]interface{}{"status": false, "message": "Todo unsuccessfully saved.", "error": "Forbidden"}
			}
	 	} else {
			t.Data["json"] = map[string]interface{}{"status": false, "message": "Todo unsuccessfully saved.", "error": "Not signed in or expired token"}
	 	}
	} else {
		t.Data["json"] = map[string]interface{}{"status": false, "message": "Todo unsuccessfully saved.", "error": "Token empty"}
	}
	t.ServeJson()
}

func (t *TodoController) GetAll() {
	result, err := models.GetAllTodo()
	if err == nil {
		t.Data["json"]=map[string]interface{}{"description": "Todo List","result": result}
	} else {
		t.Data["json"]=map[string]interface{}{"message": "Todo unsuccessfully displayed"}
	}
	t.ServeJson()
}

func (t *TodoController) Get(){
	id := t.GetString(":id")
	uid, _ := strconv.Atoi(id)
	result, err := models.GetTodo(uid)
	if err == nil {
		t.Data["json"]=map[string]interface{}{"description": "Todo details","result": result}
	} else {
		t.Data["json"]=map[string]interface{}{"message": "Failed to display todo details","result": result}
	}
	t.ServeJson()
}

func (t *TodoController) SearchByCondition() {
	Id, _ := strconv.Atoi(t.GetString(":id"))
	todo, err := models.FindTodo(Id, t.GetString(":condition"))
	if err == nil {
		t.Data["json"]=map[string]interface{}{"message": "List of Todo", "result": todo}
	} else {
		t.Data["json"]=map[string]interface{}{"message": "Todo unsuccessfully displayed"}
	}
	t.ServeJson()
}

func (t *TodoController) SearchByDatetime() {
	StartTime := t.GetString("StartTime")
	if StartTime !=""{
		EndTime := t.GetString("EndTime")
		todo, err := models.FindTodoTime(EndTime, StartTime)
		if err == nil {
			t.Data["json"]=map[string]interface{}{"message": "List of Todo", "result": todo}
		} else {
			t.Data["json"]=map[string]interface{}{"message": "Todo unsuccessfully displayed", "error": err}
		}
	} else {
		t.Data["json"]=map[string]interface{}{"message": "Todo unsuccessfully displayed", "error": "StartTime parameter is empty"}
	}
	t.ServeJson()
}

func (t *TodoController) Delete(){
	Token := t.GetString("Token")
	Id, _ := strconv.Atoi(t.GetString(":id"))
	if Id > 0 && len(Token)>0 {
		_, Level, Status := models.IsSignedIn(Token)
		if Status {
			if Level == 3 {
				result, err :=  models.DeleteTodo(Id)
				if err != nil{
					t.Data["json"]=map[string]interface{}{"Status": false,"message": err}
				} else {
					t.Data["json"]=map[string]interface{}{"Status": true,"message": result}
				}
			} else {
				t.Data["json"]=map[string]interface{}{"Status": false,"message": "Forbidden"}
			}
		} else {
			t.Data["json"]=map[string]interface{}{"Status": false,"message": "Not signed in or expired token"}
		}
	} else {
		t.Data["json"]=map[string]interface{}{"Status": false,"message": "Please check ID and token"}
	}
	t.ServeJson()
}