package cepingnet

import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type ceping_net_task_result struct {
	Cweid       int    `json:"cweid"`
	Description string `json:"description"`
	IsDeleted   []byte `json:"is_deleted"`
	IsIgnore    []byte `json:"is_ignore"`
	ItemKey     string `json:"item_key"`
	ItemValue   string `json:"item_value"`
	Level       string `json:"level"`
	Pluginid    string `json:"pluginid"`
	Points      int    `json:"points"`
	Reference   string `json:"reference"`
	ResCode     int    `json:"res_code"`
	RiskNum     int    `json:"risk_num"`
	Solution    string `json:"solution"`
	TaskId      int    `json:"task_id"`
	UpdateTime  string `json:"update_time"`
	UserId      int    `json:"user_id"`
	Wascid      int    `json:"wascid"`
}

type Unzip struct {
	Solution    string                   `json:"solution"`
	DetailInfos []map[string]interface{} `bson:"detailInfos"`
	Reference   string                   `bson:"reference"`
	TaskId      int                      `bson:"task_id"`
	Pluginid    string                   `bson:"pluginid"`
	Alert       string                   `bson:"alert"`
	ResultCode  int                      `bson:"resultCode"`
	RiskNum     int                      `bson:"risk_num"`
	Description string                   `bson:"description"`
	ItemKey     string                   `bson:"itemKey"`
	Level       string                   `bson:"level"`
}

type apptwd struct {
	TaskUid              string      `bson:"task_uid"`
	Ability              []string    `bson:"ability"`
	BackendFeedback      interface{} `bson:"backend_feedback"`
	BackendFeedbackUnzip []Unzip     `bson:"backend_feedback_unzip"`
	Code                 int         `bson:"code"`
	Description          string      `bson:"description"`
	Duration             int         `bson:"duration"`
	DurationUnit         string      `bson:"duration_unit"`
	ErrorInfo            string      `bson:"error_info"`
	RequestId            string      `bson:"request_id"`
	Sign                 string      `bson:"sign"`
	Timestamp            int         `bson:"timestamp"`
	AppraisalData        string      `bson:"appraisal_data"`
}

func Nettaskresult(taskid int, taskuid string, requestid string, ssptaskid int) {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	//ceping_ipa_task     任务id
	var ceping_net_task1 []ceping_net_task
	err = db.Table("ceping_net_task").Find(&ceping_net_task1).Error
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
	var ceping_net_task_result1 []ceping_net_task_result
	err = db.Table("ceping_net_task_result").Where("task_id = ?", taskid).Find(&ceping_net_task_result1).Error
	if err != nil {
		fmt.Println(err)
	}
	var apptwd1 apptwd

	i := 0
	score := 100
	for _, i3 := range ceping_net_task_result1 {
		apptwd1.TaskUid = taskuid
		var unzip Unzip
		_ = json.Unmarshal([]byte(i3.ItemValue), &unzip.DetailInfos)
		unzip.Solution = i3.Solution
		unzip.Reference = i3.Reference
		unzip.TaskId = ssptaskid
		unzip.Pluginid = i3.Pluginid
		unzip.Alert = i3.ItemKey
		unzip.ResultCode = i3.ResCode
		unzip.RiskNum = i3.RiskNum
		unzip.Description = i3.Description
		unzip.ItemKey = i3.ItemKey
		unzip.Level = i3.Level

		apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzip)
		apptwd1.RequestId = requestid

		if i3.Level == "Low" {
			score -= 2
		} else if i3.Level == "Medium" {
			score -= 3
		} else if i3.Level == "High" {
			score -= 4
		} else if i3.Level == "Informational" {
			score -= 1
		}

		i++
	}

	apptwd1.Ability = append(apptwd1.Ability, "web_scan")
	c := session.DB("ssp").C("tws_data")
	err = c.Insert(&apptwd1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
	//将score存入数据库
	c2 := session.DB("ssp").C("tws_task")
	//_, err := col.UpdateOne(context.Background(), bson.M{"task_uid": uid}, bson.M{"$set": bson.M{"state": "超时"}})
	_ = c2.Update(bson.M{"task_uid": taskuid}, bson.M{"$set": bson.M{"score": score}})

}
