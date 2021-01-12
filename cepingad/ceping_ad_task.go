package cepingad

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

type ceping_ad_task struct {
	Id int `json:"id"`
	ApkMd5 string `json:"apk_md5"`
	Apkshieldtype string `json:"apkshieldtype"`
	AppFilename string `json:"app_filename"`
	AppIcon string `json:"app_icon"`
	AppName string `json:"app_name"`
	AppPackagename string `json:"app_packagename"`
	AppSize int `json:"app_size"`
	AppVersionCode int `json:"app_version_code"`
	AppVersionName string `json:"app_version_name"`
	CallbackUrl string `json:"callback_url"`
	CertMd5 string `json:"cert_md5"`
	Datasource int `json:"datasource"`
	Edited []byte `json:"edited"`
	EngineIp string `json:"engine_ip"`
	FinishedItemCount int `json:"finished_item_count"`
	IsDeleted []byte `json:"is_deleted"`
	IsVersionTarget []byte `json:"is_version_target"`
	ItemKeys string `json:"item_keys"`
	MinSdkVersion int `json:"min_sdk_version"`
	ResCode int `json:"res_code"`
	SdkInfos string `json:"sdk_infos"`
	Status int `json:"status"`
	TCommit          string `json:"t_commit"`
	TUpdate          string `json:"t_update"`
	TargetSdkVersion int    `json:"target_sdk_version"`
	TotalItemCount   int    `json:"total_item_count"`
	UserId           int    `json:"user_id"`
	UsingEmulator    string `json:"using_emulator"`
	CheckType        string `json:"check_type"`
	ManualItemCount  int    `json:"manual_item_count"`
	SampleCode       int    `json:"sample_code"`
	Seting           string `json:"seting"`
	ShieldDumpCode   int    `json:"shield_dump_code"`
	TaskCode string `json:"task_code"`
	TaskTimes int `json:"task_times"`
	RequestId string `json:"request_id"`
	FailMsg string `json:"fail_msg"`
}


type CPTaskMeta struct {
	//Base
	STQueueNo        int64  `json:"st_queue_no"` //工具集排队位数
	SPPTaskID        int64  `bson:"spp_task_id"` //int的id
	RequestID        string `bson:"request_id"`  //工具集请求id
	ClientKey        string `bson:"client_key"`
	ItemCount        int64  `bson:"item_count"`         //检测项总数
	Name             string `json:"name"`               //任务名称
	CommonName       string `bson:"common_name"`        //共用名称
	Type             string `json:"type"`               //类型 (应用加固（app） sdk加固（sdk) so加固  html5加固 ）
	SDKType          string `bson:"sdk_type"`           //SDK类型
	TaskTemplateID   uint   `bson:"task_template_id"`   //任务模板id
	TaskTemplateUID  string `bson:"task_template_uid"`  //任务模板uid
	File             string `json:"file"`               //文件路径
	FileName         string `bson:"file_name"`          //原始文件名
	FileSize         int64  `bson:"file_size"`          //文件大小
	Version          string `json:"version"`            //通用版本
	WebURL           string `bson:"web_url"`            //web地址
	Score            int64  `json:"score"`              //评分
	Result           string `json:"result"`             //结果
	Summary          string `json:"summary"`            //概要
	ErrorInfo        string `bson:"error_info"`         //错误信息
	ReportStatus     int64  `bson:"report_status"`      //报告状态
	ReportLang       string `bson:"report_lang"`        //报告语言
	FinishedTime     string `bson:"finished_time"`      //完成时间
	ClientName       string `bson:"client_name"`        // 客户名称
	CreatedBy        uint   `bson:"created_by"`         // 创建人
	CreatedAt        string `bson:"created_at"`         // 创建时间
	WorkerID         uint   `bson:"worker_id"`          // 工单执行人
	Worker           string `json:"worker"`             // 工单执行人
	Icon             string `json:"icon"`               // 图标[加固]
	ReportUrl        string `bson:"report_url"`         //报告地址
	CallbackUrl      string `bson:"callback_url"`       //回调地址
	SspTaskId        uint   `bson:"ssp_task_id"`        //task_id  int类型
	TaskTemplateName string `bson:"task_template_name"` //任务模板名字
	WorkUID          string `bson:"work_uid"`           // work uid
	FileSh256        string `bson:"file_sh256"`         //上传文件的sh256
	FileMd5          string `json:"file_md5" bson:"file_md5"`           //上传文件的md5
	PkgName          string `bson:"pkg_name"`
	CertMd5          string `bson:"cert_md_5"`
	Datasource       int `json:"datasource"`
	Edited           int `json:"edited"`
	EngineIp         string `bson:"engine_ip"`
	FinishedItemCount int `bson:"finished_item_count"`
	IsDeleted int `bson:"is_deleted"`
	IsVersionTarget int `bson:"is_version_target"`
	ItemKeys string `bson:"item_keys"`
	MinSdkVersion int `bson:"min_sdk_version"`
	ResCode int `bson:"res_code"`
	SdkInfos string `bson:"sdk_infos"`
	Status int `json:"status"`
	//CreatedAt string `json:"created_at"`
	UpdatedAt string `bson:"updated_at"`
	TargetSdkVersion int `bson:"target_sdk_version"`
	//ItemCount int `json:"item_count"`
	UsingEmulator string `bson:"using_emulator"`
	CheckType string `bson:"check_type"`
	ManualItemCount int `bson:"manual_item_count"`
	TaskUid string `bson:"task_uid"`
	Key string `json:"key"`
	State string `json:"state"`
	ProductKey string `bson:"product_key"`
}


func Adtask() {

	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var ceping_ad_task1 []ceping_ad_task
	err = db.Table("ceping_ad_task").Find(&ceping_ad_task1).Error
	if err != nil {
		fmt.Println(err)
	}


	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接


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


	for _, i2 := range ceping_ad_task1 {
		fmt.Println(i2.ItemKeys)
		var ce CPTaskMeta
		ce.FileMd5 = i2.ApkMd5
		ce.FileName = i2.AppFilename
		ce.Icon = i2.AppIcon
		ce.CommonName = i2.AppName
		ce.PkgName = i2.AppPackagename
		ce.FileSize = int64(i2.AppSize)
        ce.Version = i2.AppVersionName
        ce.CallbackUrl = i2.CallbackUrl
        ce.CertMd5 = i2.CertMd5
        ce.Datasource = i2.Datasource
        ce.Edited = int(i2.Edited[0])
        ce.EngineIp = i2.EngineIp
        ce.FinishedItemCount = i2.FinishedItemCount
        ce.IsDeleted = int(i2.IsDeleted[0])
        ce.IsVersionTarget = int(i2.IsVersionTarget[0])
        ce.ItemKeys = i2.ItemKeys
        ce.MinSdkVersion = i2.MinSdkVersion
        ce.ResCode = i2.ResCode
        ce.SdkInfos = i2.SdkInfos
        ce.Status = i2.Status
        ce.CreatedAt = i2.TCommit
        ce.UpdatedAt = i2.TUpdate
        ce.TargetSdkVersion = i2.TargetSdkVersion
        ce.ItemCount = int64(i2.TotalItemCount)
        ce.CreatedBy = uint(i2.UserId)
        ce.UsingEmulator = i2.UsingEmulator
        ce.CheckType = i2.CheckType
        fmt.Println(i2.Id)
		r := rand.New(rand.NewSource(time.Now().Unix()))
	//	fmt.Println(r.Int63())
        ce.SspTaskId = uint(r.Int63())
        ce.ManualItemCount = i2.ManualItemCount
        ce.RequestID = i2.RequestId
        ce.Type = "Android"
        ce.ProductKey = "AndroidAudit"
		u1 := uuid.Must(uuid.NewV4(), nil)
		result := u1.String()
		result = strings.ReplaceAll(result, "-", "")
	//	return result
	//  ce.TaskUid = strconv.Itoa(i2.Id)
		ce.TaskUid = result
        ce.Key = "测评-安卓"
        ce.ErrorInfo = i2.FailMsg
        if i2.FailMsg == "" {
        	ce.State = "已完成"
		} else {
			ce.State = "失败"
		}

	//	i2.UserId
		var cepinguser struct {
			ContactEmail string `json:"contact_email"`
			Password  string 	`json:"password"`
			Username  string    `json:"username"`
		}

		var client_user struct{
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"upadted_at"`
			ClientKey string `json:"client_key"`
			ClientName string `json:"client_name"`
			UserAccount string `json:"user_account"`
			UserRealName string `json:"user_real_name"`
			Status int `json:"status"`
			RoleId int `json:"role_id"`
			RoleName string `json:"role_name"`
			Email string `json:"email"`
			IsAdmin bool `json:"is_admin"`
		}
		err = db.Table("ceping_user").Where("id = ?", i2.UserId).Last(&cepinguser).Error
		if err != nil {
			fmt.Println(err)
		}

		err = db2.Table("client_user").Where("client_name = ?", cepinguser.Username).Last(&client_user).Error
		if err != nil {
			fmt.Println(err)
		}
		//client_key
		//client_name
        ce.ClientName = client_user.ClientName
        ce.ClientKey = client_user.ClientKey


		//report_url
		//report_lang
		//score
		
		if i2.ResCode == 0 {
			ce.Result = "安全"
		}
		if i2.ResCode ==2 {
			ce.Result = "不安全"
		}
		if i2.ResCode ==1 {
			ce.Result = "异常"
		}
		 

		c := session.DB("ssp").C("tws_task")
		err = c.Insert(&ce)
		if err != nil {
			log.Fatal(err)
		}

		//传入对应任务的  taskid（任务id，迁移的表）   taskuid（自己生成的）
		//把对应任务的回调数据存入mongodb
		fmt.Println(23443)
		Taskresult1(i2.Id, result, i2.RequestId, i2.FailMsg)


		break
	}


}




