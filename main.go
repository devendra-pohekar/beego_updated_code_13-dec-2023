package main

import (
	"crud/controllers"
	"crud/routers"

	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	languageLablesFunc := controllers.UserController{}
	languageLablesFunc.FetchAllAndWriteInINIFiles()
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	routers.RoutersFunction()

	beego.Run()

}

// func SendOtp(c context.Context) error {
// 	email := "devendrapohekar.siliconithub@gmail.com"
// 	userName := "Devendra Pohekar"
// 	helpers.SendOTpOnMail(email, userName)

// 	return nil
// }

// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"testLog","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
// task.CreateTask("sendEmail", "0/30 * * * * *  ", SendOtp)
// req := httplib.Post("http://localhost:8080/v1/user/login")
// req.Header("Content-Type", "application/json")
// // req.Param("email", "devendrapohekar.siliconithub@gmail.com")
// // req.Param("password", "Dev@123")
// user_email := "devendrapohekar.siliconithub@gmail.com"
// user_password := "Dev@123"
// req.Body(`{"email":` + user_email + ` ,"password":` + user_password + `}`)
// res, _ := req.String()
// log.Print(res, "httplib response -----------------------")
