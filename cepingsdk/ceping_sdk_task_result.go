package cepingsdk

import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type ceping_sdk_task_result struct {
	ExceptionNum int    `json:"exception_num"`
	IsDeleted    []byte `json:"is_deleted"`
	IsIgnore     []byte `json:"is_ignore"`
	ItemKey      string `json:"item_key"`
	ItemValue    string `json:"item_value"`
	ResCode      int    `json:"res_code"`
	RiskNum      int    `json:"risk_num"`
	TaskId       int    `json:"task_id"`
	UpdateTime   string `json:"update_time"`
	UserId       int    `json:"user_id"`
}

type Unzip struct {
	TaskId         int    `bson:"task_id"`
	DetailInfos    int    `bson:"detailInfos"`
	ItemKey        string `bson:"itemKey"`
	ResultCode     int    `bson:"result_code"`
	RiskNum        int    `bson:"risk_num"`
	StartCheckDate int    `json:"startCheckDate"`
}

type apptwd struct {
	TaskUid              string                   `bson:"task_uid"`
	Ability              []string                 `bson:"ability"`
	BackendFeedback      string                   `bson:"backend_feedback"`
	BackendFeedbackUnzip []map[string]interface{} `bson:"backend_feedback_unzip"`
	Code                 int                      `bson:"code"`
	Description          string                   `bson:"description"`
	Duration             int                      `bson:"duration"`
	BurationUnit         string                   `json:"buration_unit"`
	ErrorInfo            string                   `bson:"error_info"`
	RequestId            string                   `bson:"request_id"`
	Sign                 string                   `bson:"sign"`
	TimesTamp            int                      `bson:"timestamp"`
	AppraisalData        string                   `bson:"appraisal_data"`
}

func Sdktaskresult(taskid int, taskuid string, requestid string, ssptaskid int) {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ipa_task     任务id
	var ceping_sdk_task1 []ceping_sdk_task
	err = db.Table("ceping_sdk_task").Find(&ceping_sdk_task1).Error
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
	var ceping_sdk_task_result1 []ceping_sdk_task_result
	err = db.Table("ceping_sdk_task_result").Where("task_id = ?", taskid).Find(&ceping_sdk_task_result1).Error
	if err != nil {
		fmt.Println(err)
	}
	var apptwd1 apptwd

	i := 0
	for _, i3 := range ceping_sdk_task_result1 {
		apptwd1.TaskUid = taskuid
		//	var unzip Unzip
		zzip := make(map[string]interface{})
		_ = json.Unmarshal([]byte(i3.ItemValue), &zzip)
		apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)
		apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, zzip)
		i++
	}
	apptwd1.TaskUid = taskuid
	apptwd1.RequestId = requestid
	c := session.DB("ssp").C("tws_data")
	err = c.Insert(&apptwd1)
	if err != nil {
		log.Fatal(err)
	}
}
