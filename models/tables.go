package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"

	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=demo_beego sslmode=disable")
	orm.RegisterModel(new(UserMasterTable), new(HomePagesSettingTable), new(LanguageLableLang), new(LanguageLable))
	// orm.RunSyncdb("default", false, true)

}

type UserMasterTable struct {
	UserId      int    `orm:"auto"`
	FirstName   string `orm:"size(255)"`
	LastName    string `orm:"size(255)"`
	Email       string `orm:"size(255)"`
	Password    string `orm:"size(255)"`
	Mobile      string `orm:"size(255)"`
	IsVerified  int
	OtpCode     string    `orm:size(255)`
	CreatedDate time.Time `orm:"type(datetime)"`
}

type HomePagesSettingTable struct {
	PageSettingId int `orm:"auto"`
	Section       string
	DataType      string `orm:"size(255)"`
	UniqueCode    string
	SettingData   string `orm:"type(text)"`
	CreatedDate   time.Time
	UpdatedDate   time.Time `orm:"null"`
	CreatedBy     int
	UpdatedBy     int
}

type LanguageLableLang struct {
	LangId        int `orm:"auto"`
	LanguageCode  string
	LanguageValue string
	LableCode     string `orm:"unique"`
	Section       string
}

type LanguageLable struct {
	LableId       int `orm:"auto"`
	LableCode     string
	LanguageValue string
	LanguageCode  string
	LangId        int
	Section       string
}
