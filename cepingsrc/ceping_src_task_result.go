package cepingsrc

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type ceping_src_task_result struct {
	Category   string `json:"category"`
	IsDeleted  []byte `json:"is_deleted"`
	IsIgnore   []byte `json:"is_ignore"`
	ItemKey    string `json:"item_key"`
	ItemValue  string `json:"item_value"`
	Level      string `json:"level"`
	Points     int    `json:"points"`
	ResCode    int    `json:"res_code"`
	RiskNum    int    `json:"risk_num"`
	TaskId     int    `json:"task_id"`
	UpdateTime string `json:"update_time"`
	UserId     int    `json:"user_id"`
}

type unzip struct {
	Category   string                   `bson:"category"`
	IsDeleted  []byte                   `bson:"is_deleted"`
	IsIgnore   []byte                   `bson:"is_ignore"`
	ItemKey    string                   `bson:"item_key"`
	ItemValue  string `bson:"item_value"`
	Level      string                   `bson:"level"`
	Points     int                      `bson:"points"`
	ResCode    int                      `bson:"res_code"`
	RiskNum    int                      `bson:"risk_num"`
	TaskId     int                      `bson:"task_id"`
	UpdateTime string                   `bson:"update_time"`
}

type apptwd struct {
	TaskUid              string   `bson:"task_uid"`
	Ability              []string `bson:"ability"`
	BackendFeedback      string   `json:"backend_feedback"`
	BackendFeedbackUnzip []unzip  `bson:"backend_feedback_unzip"`
	Code                 int      `bson:"code"`
	Description          string   `bson:"description"`
	Duration             int      `bson:"duration"`
	BurationUnit         string   `json:"buration_unit"`
	ErrorInfo            string   `bson:"error_info"`
	RequestId            string   `bson:"request_id"`
	Sign                 string   `bson:"sign"`
	TimesTamp            int      `bson:"timestamp"`
	AppraisalData        string   `bson:"appraisal_data"`
}

func Srctaskresult(taskid int, taskuid string, requestid string, ssptaskid int) {
	db5, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db5.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式

	db2	, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接


	//ceping_ipa_task     任务id
	var ceping_src_task1 []ceping_src_task
	err = db5.Table("ceping_src_task").Find(&ceping_src_task1).Error
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

	db6, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db6.Close() //用完之后关闭数据库连接

	//i2.Id  任务id
	var ceping_src_task_result1 []ceping_src_task_result
	err = db6.Table("ceping_src_task_result").Where("task_id = ?", taskid).Find(&ceping_src_task_result1).Error
	if err != nil {
		fmt.Println(err)
	}
	var apptwd1 apptwd

	i := 0
	for _, i3 := range ceping_src_task_result1 {
		apptwd1.TaskUid = taskuid
		apptwd1.Ability = append(apptwd1.Ability, i3.ItemKey)

		var unzi unzip
		unzi.Level = i3.Level
		unzi.ItemKey = i3.ItemKey
		unzi.RiskNum = i3.RiskNum
		unzi.ResCode = i3.ResCode
		unzi.TaskId = ssptaskid
		unzi.Category = i3.Category
		unzi.IsDeleted = i3.IsDeleted
		unzi.IsIgnore = i3.IsIgnore
		unzi.Points = i3.Points
		unzi.UpdateTime = i3.UpdateTime

		/*var itemv []map[string]interface{}
		for i2, i4 := range i3.ItemValue {

		}*/
     //   fmt.Println(i3.ItemValue)
		unzi.ItemValue = i3.ItemValue
		apptwd1.BackendFeedbackUnzip = append(apptwd1.BackendFeedbackUnzip, unzi)
		i++
	}
	apptwd1.TaskUid = taskuid
	apptwd1.RequestId = requestid
	c := session.DB("ssp").C("tws_data")
	err = c.Insert(&apptwd1)
	if err != nil {
		log.Fatal(err)
	}

	db4, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db4.Close() //用完之后关闭数据库连接
	var score int
	err = db4.Raw(`select  FLOOR((1 - sum(case when t.item_value is null then 0 else i.score end) /
    (case when sum(i.score) <= 0 then 1 else sum(i.score) end)) * 100) 
    from ceping_src_task_result t right outer join ceping_src_audit_item i
    on t.item_key = i.rule_key and t.task_id = ?`, taskid).Row().Scan(&score)
	/*err = db4.Raw(`select  FLOOR((1 - sum(case when t.item_value is null then 0 else i.score end) /
					(case when sum(i.score) <= 0 then 1 else sum(i.score) end)) * 100) 
					from ceping_src_task_result t right outer join ceping_src_audit_item_cn i 
					on t.item_key = i.rule_key and t.task_uid = ?`, taskuid).Row().
		Scan(&score)*/
	fmt.Println(score)
	/*err = db.Raw(`select  FLOOR((1 - sum(case when t.item_value is null then 0 else i.score end) /
    (case when sum(i.score) <= 0 then 1 else sum(i.score) end)) * 100) 
    from ceping_src_task_result t right outer join ceping_src_audit_item i
    on t.item_key = i.rule_key and t.task_id = ?`, taskid).Row().Scan(&score)*/

	c2 := session.DB("ssp").C("tws_task")
	//_, err := col.UpdateOne(context.Background(), bson.M{"task_uid": uid}, bson.M{"$set": bson.M{"state": "超时"}})
	_ = c2.Update(bson.M{"task_uid": taskuid}, bson.M{"$set": bson.M{"score": score}})

}
