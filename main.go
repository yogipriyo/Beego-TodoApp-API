package main

import (
	_ "github.com/yogipriyo/todo_app3/routers"
	models "github.com/yogipriyo/todo_app3/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func init() {
    orm.RegisterDriver("postgres", orm.DR_Postgres)
    orm.RegisterDataBase("default", "postgres", "user=postgres password=09030015 dbname=todo_app sslmode=disable")
    orm.RegisterModel(new(models.Users), new(models.Level), new(models.Todo), new(models.Category), new(models.Status))

    // Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := false

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
    //db, err := sql.Open("postgres", "user=postgres password=09030015 dbname=online-bus-api sslmode=disable")

    orm.Debug = false
}


func main() {
	beego.Run()
}

