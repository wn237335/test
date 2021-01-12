package cepingad

import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
	"math"

	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type ceping_ad_task_result struct {
	EditDesc     string `bson:"edit_desc"`
	EditDetail   string `bson:"edit_detail"`
	ExceptionNum int    `bson:"exception_num"`
	IsDeleted    []byte `bson:"is_deleted"`
	IsIgnore     []byte `bson:"is_ignore"`
	ItemKey      string `bson:"item_key"`
	ItemValue    string `bson:"item_value"`
	OtherPoints  int    `bson:"other_points"`
	OtherRiskNum int    `bson:"other_risk_num"`
	OtherSdk     string `bson:"other_sdk"`
	OtherValue   string `bson:"other_value"`
	Points       string `json:"points"`
	ResCode      int    `bson:"res_code"`
	ResVal       string `bson:"res_val"`
	RiskNum      int    `bson:"risk_num"`
	TaskId       int    `bson:"task_id"`
	UpdateTime   string `bson:"update_time"`
	UserId       string `bson:"user_id"`
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
	MinSdkVersion    int    `json:"minSdkVersion"`
	TargetSdkVersion int    `json:"targetSdkVersion"`
	AppMd5           string `json:"appMd5"`
	AppName          string `json:"appName"`
	AppPackageName   string `json:"appPackageName"`
	AppVersion       string `json:"appVersion"`
	AppSize          int    `json:"appSize"`
	ApkSign          struct {
		Issuer   string `json:"issuer"`
		Serial   string `json:"serial"`
		SignInfo string `json:"signInfo"`
		Subject  string `json:"subject"`
	} `json:"apkSign"`
}

type apptwd struct {
	TaskUid string `bson:"task_uid"`
	AppInfo struct {
		Target int `json:"target"`
		Min    int `json:"min"`
		Sign   string `json:"sign"`
		PkgName string `bson:"pkg_name"`
		AppName string `bson:"app_name"`
		Version string `json:"version"`
		Size    int    `json:"size"`
		Md5     string `json:"md5"`
	} `bson:"app_info"`
	AppType         string   `bson:"app_type"`
	Ability         []string `json:"ability"`
	BackendFeedback string   `bson:"backend_feedback"`

//	BackendFeedbackUnzip []Unzip `json:"backend_feedback_unzip"`
	BackendFeedbackUnzip []map[string]interface{} `json:"backend_feedback_unzip"`

	Code          int    `json:"code"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	DurationUnit  string `bson:"duration_unit"`
	ErrorInfo     string `bson:"error_info"`
	ItemCount     int    `bson:"item_count"`
	RequestId     string `bson:"request_id"`
	Sign          string `json:"sign"`
	TimeStamp     int    `bson:"time_stamp"`
	AppraisalData string
}

func Taskresult1(taskid int, taskuid string, requestid string, failmsg string) {

	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ad_task     任务id
	var ceping_ad_task1 []ceping_ad_task
	err = db.Table("ceping_ad_task").Find(&ceping_ad_task1).Error
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
	var ceping_ad_task_result1 []ceping_ad_task_result
	err = db.Table("ceping_ad_task_result").Where("task_id = ?", taskid).Find(&ceping_ad_task_result1).Error
	if err != nil {
		fmt.Println(err)
	}
	var apptwd1 apptwd

	i := 0
	var securityScore, AllScore int
	for _, i3 := range ceping_ad_task_result1 {
		apptwd1.TaskUid = taskuid
		apptwd1.AppType = "Android"
		apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)
		str := "ok"
		if i3.ItemKey == "sec_infos" {
			appinfo1 := appinfo{}
			_ = json.Unmarshal([]byte(i3.ItemValue), &appinfo1)
			apptwd1.AppInfo.Sign = fmt.Sprintf("所有者:%s\n发布者:%s", appinfo1.ApkSign.Subject, appinfo1.ApkSign.Issuer)
		    apptwd1.AppInfo.Target = appinfo1.TargetSdkVersion
		    apptwd1.AppInfo.Min = appinfo1.MinSdkVersion
		    apptwd1.AppInfo.PkgName = appinfo1.AppPackageName
		    apptwd1.AppInfo.AppName = appinfo1.AppName
		    apptwd1.AppInfo.Version = appinfo1.AppVersion
		    apptwd1.AppInfo.Size = appinfo1.AppSize
		    apptwd1.AppInfo.Md5 = appinfo1.AppMd5
			str = "notok"
		}

		if str == "ok" {
		//	var unzip Unzip
			unzz := make(map[string]interface{})
			_ = json.Unmarshal([]byte(i3.ItemValue), &unzz)

			if int(unzz["resultCode"].(float64)) < 300 {
				AllScore += int(unzz["resultCode"].(float64))
			}
			if int(unzz["resultCode"].(float64)) > 100 && int(unzz["resultCode"].(float64)) < 200 {
				securityScore += int(unzz["resultCode"].(float64))
			}

			apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzz)
		//	_ = json.Unmarshal([]byte(i3.ItemValue), &unzip)
		//	apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzip)
		}
		apptwd1.RequestId = requestid
		i++
	}
	apptwd1.ErrorInfo = failmsg
	apptwd1.ItemCount = i
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
	//_, err := col.UpdateOne(context.Background(), bson.M{"task_uid": uid}, bson.M{"$set": bson.M{"state": "超时"}})
	_ = c2.Update(bson.M{"task_uid": taskuid}, bson.M{"$set": bson.M{"score": score}})
}

/*func Taskresult() {

	db, err := gorm.Open("mysql", "root:www@admin@2020@(172.16.42.66:33060)/securityte_java?charset=utf8mb4&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ad_task     任务id
	var ceping_ad_task1 []ceping_ad_task
	err = db.Table("ceping_ad_task").Find(&ceping_ad_task1).Error
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

	for _, i2 := range ceping_ad_task1 {
		//i2.Id  任务id
		var ceping_ad_task_result1 []ceping_ad_task_result
		err = db.Table("ceping_ad_task_result").Where("task_id = ?", i2.Id).Find(&ceping_ad_task_result1).Error
		if err != nil {
			fmt.Println(err)
		}
		var apptwd1 apptwd
		fmt.Println(i2.AppVersionName)
		apptwd1.AppInfo.Target = i2.TargetSdkVersion
		apptwd1.AppInfo.Min = i2.MinSdkVersion
		apptwd1.AppInfo.Version = i2.AppVersionName
		apptwd1.AppInfo.PkgName = i2.AppPackagename
		apptwd1.AppInfo.AppName = i2.AppName
		apptwd1.AppInfo.Md5 = i2.ApkMd5
		apptwd1.AppInfo.Size = i2.AppSize

		i := 0
		for _, i3 := range ceping_ad_task_result1 {
			apptwd1.TaskUid = strconv.Itoa(i3.TaskId)
			apptwd1.AppType = "Android"
			apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)
			var unzip Unzip
			_ = json.Unmarshal([]byte(i3.ItemValue), &unzip)
			apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzip)
			//	apptwd1.BackendFeedbackUnzip.Packer
			//	apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip,i3.ItemValue )
			apptwd1.RequestId = i2.RequestId
			i++
		}
		apptwd1.ItemCount = i
		c := session.DB("ssp").C("tws_data")
		err = c.Insert(&apptwd1)
		if err != nil {
			log.Fatal(err)
		}
	}

}*/
