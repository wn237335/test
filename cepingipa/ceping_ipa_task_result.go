package cepingipa

import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math"
	"time"
)

type ceping_ipa_task_result struct {
	EditDesc     string `json:"edit_desc"`
	EditDetail   string `json:"edit_detail"`
	ExceptionNum string `json:"exception_num"`
	IsDeleted    []byte `json:"is_deleted"`
	IsIgnore     []byte `json:"is_ignore"`
	ItemKey      string `json:"item_key"`
	ItemValue    string `json:"item_value"`
	Points       int    `json:"points"`
	ResCode      int    `json:"res_code"`
	ResVal       string `json:"res_val"`
	RiskNum      int    `json:"risk_num"`
	TaskId       int    `json:"task_id"`
	UpadteTime   string `json:"upadte_time"`
	UserId       int    `json:"user_id"`
}

type Unzip struct {
	UseSeconds     int                      `bson:"useSeconds"`
	StartCheckDate int                      `bson:"startCheckDate"`
	TaskId         int                      `json:"taskId"`
	ResultCode     int                      `json:"resultCode"`
	Packer         string                   `json:"packer"`
	TypeInfos      []map[string]interface{} `json:"typeInfos"`
	DetailInfos    []map[string]interface{} `json:"detailInfos"`
	ItemKey        string                   `json:"itemKey"`
	RiskNum        int                      `json:"riskNum"`
}

type appinfo struct {
	AppMd5         string `json:"appMd5"`
	AppName        string `json:"appName"`
	AppPackageName string `json:"appPackageName"`
	AppVersion     string `json:"appVersion"`
	AppSize        int    `json:"appSize"`
	SignInfo       struct {
		ApplicationIdentifier string `json:"applicationIdentifier"`
		TeamIdentifier        string `json:"teamIdentifier"`
		TeamName              string `json:"teamName"`
		Uuid                  string `json:"uuid"`
	} `json:"signInfo"`
}

type apptwd struct {
	TaskUid string `bson:"task_uid"`
	AppInfo struct {
		Sign struct {
			ApplicationIdentifier string `json:"application_identifier"`
			TeamIdentifier        string `json:"team_identifier"`
			TeamName              string `json:"team_name"`
			Uuid                  string `json:"UUID"`
		} `json:"sign"`
		PkgName string `bson:"pkg_name"`
		AppName string `Bson:"app_name"`
		Version string `json:"version"`
		Size    int    `json:"size"`
		Md5     string `json:"md5"`
	} `bson:"app_info"`
	AppType              string      `bson:"app_type"`
	ItemCount            int         `bson:"item_count"`
	RequestId            string      `bson:"request_id"`
	Ability              []string    `bson:"ability"`
	BackendFeedback      interface{} `bson:"backend_feedback"`
//	BackendFeedbackUnzip []Unzip     `json:"backend_feedback_unzip"`
	BackendFeedbackUnzip []map[string]interface{}     `json:"backend_feedback_unzip"`

	Code          int    `json:"code"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	DurationUnit  string `json:"duration_unit"`
	ErrorInfo     string `json:"error_info"`
	Sign          string `json:"sign"`
	Timestamp     int    `json:"timestamp"`
	AppraisalData string `bson:"appraisal_data"`
}




func Ipataskresult1(taskid int, taskuid string, requestid string, errorinfo string) {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ipa_task     任务id
	var ceping_ipa_task1 []ceping_ipa_task
	err = db.Table("ceping_ipa_task").Find(&ceping_ipa_task1).Error
	if err != nil {
		fmt.Println(err)
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{utils.Sspmongodbserver()},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  utils.Sspmongodbname(),
		Source:    "admin",
		Username:  utils.Sspmongodbuser(),
		Password:  utils.Sspmongodbpwd(),
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//i2.Id  任务id
	var ceping_ipa_task_result1 []ceping_ipa_task_result
	err = db.Table("ceping_ipa_task_result").Where("task_id = ?", taskid).Find(&ceping_ipa_task_result1).Error
	if err != nil {
		fmt.Println(err)
	}
	var apptwd1 apptwd

	i := 0
	var securityScore, AllScore int
	for _, i3 := range ceping_ipa_task_result1 {
		apptwd1.TaskUid = taskuid
		apptwd1.AppType = "IOS"
		apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)
		str := "ok"
		if i3.ItemKey == "ios_sec_infos" {
			appinfo1 := appinfo{}
			_ = json.Unmarshal([]byte(i3.ItemValue), &appinfo1)
			apptwd1.AppInfo.Version = appinfo1.AppVersion
			apptwd1.AppInfo.PkgName = appinfo1.AppPackageName
			apptwd1.AppInfo.AppName = appinfo1.AppName
			apptwd1.AppInfo.Md5 = appinfo1.AppMd5
			apptwd1.AppInfo.Size = appinfo1.AppSize
			apptwd1.AppInfo.Sign.ApplicationIdentifier = appinfo1.SignInfo.ApplicationIdentifier
			apptwd1.AppInfo.Sign.TeamIdentifier = appinfo1.SignInfo.TeamIdentifier
			apptwd1.AppInfo.Sign.TeamName = appinfo1.SignInfo.TeamName
			apptwd1.AppInfo.Sign.Uuid = appinfo1.SignInfo.Uuid
			str = "notok"
		}
		if str == "ok" {
			/*var unzip Unzip
			_ = json.Unmarshal([]byte(i3.ItemValue), &unzip)
			apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzip)*/
			unzz := make(map[string]interface{})
			_ = json.Unmarshal([]byte(i3.ItemValue), &unzz)

			if int(unzz["resultCode"].(float64)) < 300 {
				AllScore += int(unzz["resultCode"].(float64))
			}
			if int(unzz["resultCode"].(float64)) > 100 && int(unzz["resultCode"].(float64)) < 200 {
				securityScore += int(unzz["resultCode"].(float64))
			}

			apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzz)
		}
		apptwd1.RequestId = requestid
		i++
	}

	apptwd1.ErrorInfo = errorinfo
	apptwd1.ItemCount = i
	fmt.Println()
	c := session.DB("ssp").C("tws_data")
	err = c.Insert(&apptwd1)
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	if AllScore != 0 {
		score = int(math.Round(float64(securityScore) / float64(AllScore) * 100)) // 总分
	} else {
		score = 100
	}
	c2 := session.DB("ssp").C("tws_task")
	_ = c2.Update(bson.M{"task_uid": taskuid}, bson.M{"$set": bson.M{"score": score}})

}

/*func Ipataskresult(taskid int, taskuid string) {
	db, err := gorm.Open("mysql", "root:www@admin@2020@(172.16.42.66:33060)/securityte_java?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ipa_task     任务id
	var ceping_ipa_task1 []ceping_ipa_task
	err = db.Table("ceping_ipa_task").Find(&ceping_ipa_task1).Error
	if err != nil {
		fmt.Println(err)
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"172.16.102.56"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "ssp",
		Source:    "admin",
		Username:  "root",
		Password:  "Ky_Monogo_2019!",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	for _, i2 := range ceping_ipa_task1 {
		//i2.Id  任务id
		var ceping_ipa_task_result1 []ceping_ipa_task_result
		err = db.Table("ceping_ipa_task_result").Where("task_id = ?", i2.Id).Find(&ceping_ipa_task_result1).Error
		if err != nil {
			fmt.Println(err)
		}
		var apptwd1 apptwd

		i := 0
		for _, i3 := range ceping_ipa_task_result1 {
			apptwd1.TaskUid = strconv.Itoa(i3.TaskId)
			apptwd1.AppType = "IOS"
			apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)
			str := "ok"
			if i3.ItemKey == "ios_sec_infos" {
				appinfo1 := appinfo{}
				_ = json.Unmarshal([]byte(i3.ItemValue), &appinfo1)
				apptwd1.AppInfo.Version = appinfo1.AppVersion
				apptwd1.AppInfo.PkgName = appinfo1.AppPackageName
				apptwd1.AppInfo.AppName = appinfo1.AppName
				apptwd1.AppInfo.Md5 = appinfo1.AppMd5
				apptwd1.AppInfo.Size = appinfo1.AppSize
				apptwd1.AppInfo.Sign.ApplicationIdentifier = appinfo1.SignInfo.ApplicationIdentifier
				apptwd1.AppInfo.Sign.TeamIdentifier = appinfo1.SignInfo.TeamIdentifier
				apptwd1.AppInfo.Sign.TeamName = appinfo1.SignInfo.TeamName
				apptwd1.AppInfo.Sign.Uuid = appinfo1.SignInfo.Uuid
				str = "notok"
			}
			if str == "ok" {
				var unzip Unzip
				_ = json.Unmarshal([]byte(i3.ItemValue), &unzip)
				apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzip)
			}
			apptwd1.RequestId = i2.RequestId
			i++
		}
		apptwd1.ItemCount = i
		c := session.DB("ssp").C("tws_data")
		err = c.Insert(&apptwd1)
		if err != nil {
			log.Fatal(err)
		}
		break
	}

}*/
