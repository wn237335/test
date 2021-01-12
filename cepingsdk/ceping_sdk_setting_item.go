package cepingsdk
import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
)

type ceping_sdk_setting_item struct {
	ItemKey string `json:"item_key"`
	UserId  int    `json:"user_id"`
}

type user_item struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Item   string `json:"item"`
	Type   string `json:"type"` //Android Ios Sdk
}

type cepinguser struct {
	Id                        int    `json:"id"`
	AndroidStudioId           int    `json:"android_studio_id"`
	ApiSignature              string `json:"api_signature"`
	AuthType                  string `json:"auth_type"`
	CepingApps                int    `json:"ceping_apps"`
	CepingTimes               int    `json:"ceping_times"`
	Company                   string `json:"company"`
	ContactEmail              string `json:"contact_email"`
	Enable                    []byte `json:"enable"`
	IosEndtime                string `json:"ios_endtime"`
	IosStudioId               int    `json:"ios_studio_id"`
	IpaApps                   int    `json:"ipa_apps"`
	IpaTimes                  int    `json:"ipa_times"`
	IsFirstlogin              []byte `json:"is_firstlogin"`
	LastUseday                int    `json:"last_useday"`
	NetEndtime                string `json:"net_endtime"`
	NetTimes                  int    `json:"net_times"`
	Password                  string `json:"password"`
	Platform                  string `json:"platform"`
	Platforms                 string `json:"platforms"`
	ReportFormats             string `json:"report_formats"`
	ReportIosLanguage         int    `json:"report_ios_language"`
	ReportLanguage            int    `json:"report_language"`
	ReportTemplate            string `json:"report_template"`
	RoleType                  string `json:"role_type"`
	SdkApps                   int    `json:"sdk_apps"`
	SdkEndtime                string `json:"sdk_endtime"`
	SourceTimes               int    `json:"source_times"`
	Special                   string `json:"special"`
	TEndtime                  string `json:"t_endtime"`
	TRegtime                  string `json:"t_regtime"`
	TUpdatetime               string `json:"t_updatetime"`
	Username                  string `json:"username"`
	AndroidReportTemplate     string `json:"android_report_template"`
	IosReportTemplate         string `json:"ios_report_template"`
	SourceCount               int    `json:"source_count"`
	SourceEndTime             string `json:"source_end_time"`
	SourceStartTime           string `json:"source_start_time"`
	SourceReportTemplate      string `json:"source_report_template"`
	MiniProgramEndtime        string `json:"mini_program_endtime"`
	MiniProgramTimes          int    `json:"mini_program_times"`
	ReportMiniProgramLanguage int    `json:"report_mini_program_language"`
	MiniProgramApps           int    `json:"mini_program_apps"`
	ReportWatermark           string `json:"report_watermark"`
}

func Sdksettingitem1() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	var cepinguser []cepinguser
	err = db.Table("ceping_user").Find(&cepinguser).Error
	if err != nil {
		fmt.Println(err)
	}

	for _, i6 := range cepinguser {
		var ceping_sdk_setting_item1 []ceping_sdk_setting_item
		err = db.Table("ceping_sdk_setting_item").Where("user_id = ?", i6.Id).Find(&ceping_sdk_setting_item1).Error
		if err != nil {
			fmt.Println(err)
			continue
		}
		if ceping_sdk_setting_item1 == nil {
			continue
		}
		jso, err := ioutil.ReadFile("resources/ceping_item.json")
		if err != nil {
			fmt.Println(err)
		}

		jss := make([]interface{}, 0)
		_ = json.Unmarshal(jso, &jss)
		//			logger.Error(jss["Android"])
		result := make(map[string]interface{})
		for _, i2 := range jss {
			if i2.(map[string]interface{})["type"] == "SDK" {
				result[i2.(map[string]interface{})["item"].(string)] = 1
			}
		}

		for _, i7 := range ceping_sdk_setting_item1 {
			if result[i7.ItemKey] != nil {
				if result[i7.ItemKey] == 1 {
					delete(result, i7.ItemKey)
				}
			}
		}

		var ceping_sdk_setting_item2 user_item

		var aps_user struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			RealName string `json:"real_name"`
			Password string `json:"password"`
		}

		err = db2.Table("aps_user").Where("name = ?", i6.Username).Last(&aps_user).Error
		if err != nil {
			fmt.Println(err)
		}

		ceping_sdk_setting_item2.UserId = aps_user.Id
		re, _ := json.Marshal(result)
		ceping_sdk_setting_item2.Item = string(re)
		ceping_sdk_setting_item2.Name = i6.Username
		ceping_sdk_setting_item2.Type = "SDK"
		//将数据存入表中
		err = db2.Table("user_item").Create(&ceping_sdk_setting_item2).Error
		if err != nil {
			fmt.Println(err)
		}

	}
}
