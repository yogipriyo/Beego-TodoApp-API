package controllers

import (
	"strconv"

	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"crypto/hmac"

	"encoding/hex"
	"github.com/astaxie/beego"
	models "github.com/yogipriyo/todo_app3/models"
	// "fmt"
	// "log"
)

// oprations for user
type UserController struct {
	beego.Controller
}

//

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func genToken(rand string, spid string) string {
	key := []byte(spid)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(rand))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *UserController) Post() {
	NewPassword := GetMD5Hash(u.GetString("Password"))
	Level, _ := u.GetInt("Level")
	user := models.Users{Email: u.GetString("Email"), Fullname: u.GetString("Fullname"), MobileNumber: u.GetString("Mobile"), Gender: u.GetString("Gender"), Password: NewPassword, Level: Level}
	
	uid, err := models.AddUser(&user)
	if err == nil {
		u.Data["json"] = map[string]interface{}{"message": "User successfully saved.", "result": uid}
	}
	u.ServeJson()
}

func (u *UserController) GetAll() {
	users, err, _ := models.GetAllUsers()
	if err == nil {
		//u.Data["json"] = users
		u.Data["json"] = map[string]interface{}{"description": "Users list", "result": users}
	} 
	u.ServeJson()
}

func (u *UserController) Get() {
	id := u.GetString(":id")
	uid, _ := strconv.Atoi(id)
	if uid != 0 {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = map[string]interface{}{"description": "Users details", "result": user}
			// u.Data["json"] = user
		}
	}
	u.ServeJson()
}

func (u *UserController) Delete() {
	id := u.GetString(":id")
	uid, _ := strconv.Atoi(id)
	if uid != 0 {
		result, err := models.DeleteUser(uid)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = result
		}
	}
	u.ServeJson()
}

func (u *UserController) Login() {
	NewPassword := GetMD5Hash(u.GetString("Password"))
	result, user := models.Find(u.GetString("Email"), NewPassword)
	
	if result {
		rand := randString(32)
		token := genToken(rand, u.GetString("Email"))
		models.UpdateToken(token, user.Id)
		// u.Data["json"] = map[string]interface{}{"status": true, "message": "User successfully logged in.", "user_id": user.Id, "token": token}
		u.Data["json"] = map[string]interface{}{"message": "User successfully logged in.", "user_id": user.Id, "token": token}
	} else {
		// u.Data["json"] = map[string]interface{}{"status": false, "message": "Login failed."}
		u.Data["json"] = map[string]interface{}{"message": "Login failed."}
	}
	u.ServeJson()
}

func (u *UserController) Logout() {
	if u.GetString("Token") != "" {
		result, err := models.DestroyToken(u.GetString("Token"))
		if result {
			u.Data["json"] = map[string]interface{}{"status": true, "message": "Logout successfully"}
		} else {
			u.Data["json"] = map[string]interface{}{"status": false, "message": err}
		}
	} else {
		u.Data["json"] = map[string]interface{}{"status": false, "message": "Token is empty, please insert token"}
	}
	u.ServeJson()
}