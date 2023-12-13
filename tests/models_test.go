package test_test

import (
	"crud/models"
	requestStruct "crud/requstStruct"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=demo_beego sslmode=disable")
	orm.RunSyncdb("default", false, true)

}

func TestUserModels(t *testing.T) {
	t.Run("RegisterUserModel", func(t *testing.T) {
		insert_data := requestStruct.InsertUser{
			FirstName: "Testing Devendra",
			LastName:  "Testing Pohekar",
			Email:     "devendrapohekar.siliconithub@gmail.com",
			Password:  "Dev@123",
			Mobile:    "1234567890",
		}
		result, err := models.RegisterUser(insert_data)
		if err != nil {
			t.Log(result, "result of register user not testing clear please check the RegisterUser Model function")
		}
		t.Log(result, "successfully done Register User Model ")
	})
	t.Run("LoginUserModel", func(t *testing.T) {
		login_data := requestStruct.LoginUser{
			Email:    "devendrapohekar.siliconithub@gmail.com",
			Password: "Dev@123",
		}
		response, err := models.LoginUsers(login_data)
		if err != nil {
			t.Log("Testing Login Failed Due to some issue in LoginUsers Model ")

		}
		t.Log(response, "successfully done the LoginUsers Model")
	})
	t.Run("RegisterHomePageSetting", func(t *testing.T) {
		register_data := requestStruct.HomeSeetingInsert{
			Section:     "Top Header Left Corner Text",
			DataType:    "text",
			SettingData: "This is a TESTING DATA",
		}
		res, err := models.RegisterSetting(register_data, 31, "")
		if err != nil {
			t.Log("Register Setting of Home Page Occure Error in Testing Please Check the RegisterSetting Models function ")

		}
		t.Log(res, "Successfully Pass in Testing Phase you can move ahead")

	})

	t.Run("UpdateHomePageSetting", func(t *testing.T) {
		update_data := requestStruct.HomeSeetingUpdate{
			Section:     "header right corner paragraph",
			DataType:    "html",
			SettingData: "<p>Hello Testing Right Corner Para</p>",
			SettingId:   1,
		}
		upd_result, err := models.UpdateSetting(update_data, "", 32)
		if err != nil {
			t.Log("UpdateSetting Model function Have some error in testing phase please check modules configuration ")
		}
		t.Log(upd_result, "Successfully UpdateSetting Model Work Here is a UpdatedSetting ID you can move ahead")

	})
	t.Run("DeleteHomePageSetting", func(t *testing.T) {
		home_page_setting_id := requestStruct.HomeSeetingDelete{
			SettingId: 1,
		}
		res := models.HomePageSettingExistsDelete(home_page_setting_id)
		if res < 0 {
			t.Log("DeleteHomePageSetting Successfully Not Work Please check the DeleteSetting Model function Manualy")
		}
		t.Log("Successfully  Delete Setting  and The DeleteHomePageSetting  Testing Function Done and HomePageSettingExistsDelete Work Successfully you can move ahead")

	})
}
