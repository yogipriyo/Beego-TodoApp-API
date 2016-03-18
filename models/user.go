package models

import (
	//"errors"
	//"strconv"
	"time"
	"log"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       		int			`orm:"column(id);auto"`
	Email	 		string		`orm:"column(email)"`
	Fullname		string		`orm:"column(fullname)"`
	MobileNumber	string		`orm:"column(mobile_number)"`
	Gender			string		`orm:"column(gender)"`
	Password 		string		`orm:"column(password)"`
	Level			int			`orm:"column(level)"`
	Token			string		`orm:"column(token)"`
	TokenCreated	time.Time 	`orm:"column(token_created);type(datetime);null"`
	TokenExpired	time.Time 	`orm:"column(token_expired);type(datetime);null"`
}

type Level struct {
	Id 		int 	`orm:"column(id);auto"`
	Name	string	`orm:"column(name)"`
}

//create new user
func AddUser(t *Users) (*Users, error) {
	o := orm.NewOrm()
	_, err := o.Insert(t)
	if err != nil {
		log.Print(err)
	}
	return t, err
}

// Get All user
func GetAllUsers() ([]*Users, error, int64) {
	o := orm.NewOrm()
	var result []*Users
	num, err := o.QueryTable(Users{}).RelatedSel().All(&result)
	if err != orm.ErrNoRows && num > 0 {
		return result, nil, num
	} else {
		return nil, err, num
	}
}

// Get user
func GetUser(uid int) (*Users, error) {
	o := orm.NewOrm()
	result := &Users{Id: uid}

	err := o.Read(result)

	if err == orm.ErrNoRows {
	    return nil, err
	} else if err == orm.ErrMissPK {
	    return nil, err
	} else {
	    return result, nil
	}
}

// Delete user
func DeleteUser(uid int) (string, error) {
	o := orm.NewOrm()
	//result := &User{Id: uid}

	if _, err := o.Delete(&Users{Id: uid}); err == nil {
		return "User deleted", nil
	} else {
		return "Failed to delete user", err
	}
}


//additional func 
//find user
func Find(email, password string)(bool, Users){
	o := orm.NewOrm()
	
	var user Users
	err := o.Raw("SELECT * FROM users WHERE email = ? and password = ?", email, password).QueryRow(&user)

	//log.Println(user)
	//log.Println(err)

	if err != nil{
		return false, user
	} else {
		return true, user
	}
}

//update token
func UpdateToken(token string, id int){
	o := orm.NewOrm()
	user := &Users{Id: id}
	err := o.Read(user)
	if err == nil {
		CurrentTime := time.Now()
		user.Token = token
		user.TokenCreated = CurrentTime
		user.TokenExpired = CurrentTime.Add(time.Hour * 1)
		//user.TokenCreated = time.Now()
		if num, err := o.Update(user); err == nil{
			log.Println(num, " token updated")
		} else {
			log.Println(num, " token update failed")
		}
	}
}

func DestroyToken(token string)(result bool, err error){
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE users SET token = ? WHERE token = ?", "", token).Exec()
	if err == nil {
	    return true, err
	} else {
		return false, err
	}
}

func IsSignedIn(token string)(id int, level int, status bool){
	o := orm.NewOrm()
	user := &Users{Token: token}
	err := o.Read(user, "Token")

	CurrentTime := time.Now()
	diff:= CurrentTime.Sub(user.TokenExpired)
	//log.Println(diff)

	if err == nil && diff<0{
		return user.Id, user.Level, true
	} else {
		//log.Println(err)
		return user.Id, user.Level, false
	}
}