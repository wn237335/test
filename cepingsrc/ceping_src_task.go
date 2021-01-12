package cepingsrc

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"strings"
	"time"
)

type ceping_src_task struct {
	Id              int    `json:"id"`
	Bugs            int    `json:"bugs"`
	CallbackUrl     string `json:"callback_url"`
	CodeSmells      int    `json:"code_smells"`
	IsDeleted       []byte `json:"is_deleted"`
	IsFinished      []byte `json:"is_finished"`
	LinkTaskId      int    `json:"link_task_id"`
	LinkTaskType    string `json:"link_task_type"`
	ResCode         int    `json:"res_code"`
	ShowFinished    []byte `json:"show_finished"`
	SrcFilename     string `json:"src_filename"`
	SrcMd5          string `json:"src_md5"`
	SrcSize         int    `json:"src_size"`
	SrcVersion      string `json:"src_version"`
	TCommit         string `json:"t_commit"`
	TUpdate         string `json:"t_update"`
	TotalLine       int    `json:"total_line"`
	UserId          int    `json:"user_id"`
	Vulnerabilities int    `json:"vulnerabilities"`
	RequestId       string `json:"request_id"`
	FailMsg         string `json:"fail_msg"`
}

type CPTaskMeta struct {
	//Base
	STQueueNo         int64  `json:"st_queue_no"` //工具集排队位数
	SPPTaskID         int64  `bson:"spp_task_id"` //int的id
	RequestID         string `bson:"request_id"`  //工具集请求id
	ClientKey         string `bson:"client_key"`
	ItemCount         int64  `bson:"item_count"`               //检测项总数
	Name              string `json:"name"`                     //任务名称
	CommonName        string `bson:"common_name"`              //共用名称
	Type              string `json:"type"`                     //类型 (应用加固（app） sdk加固（sdk) so加固  html5加固 ）
	SDKType           string `bson:"sdk_type"`                 //SDK类型
	TaskTemplateID    uint   `bson:"task_template_id"`         //任务模板id
	TaskTemplateUID   string `bson:"task_template_uid"`        //任务模板uid
	File              string `json:"file"`                     //文件路径
	FileName          string `bson:"file_name"`                //原始文件名
	FileSize          int64  `bson:"file_size"`                //文件大小
	Version           string `json:"version"`                  //通用版本
	WebURL            string `bson:"web_url"`                  //web地址
	Score             int64  `json:"score"`                    //评分
	Result            string `json:"result"`                   //结果
	Summary           string `json:"summary"`                  //概要
	ErrorInfo         string `bson:"error_info"`               //错误信息
	ReportStatus      int64  `bson:"report_status"`            //报告状态
	ReportLang        string `bson:"report_lang"`              //报告语言
	FinishedTime      string `bson:"finished_time"`            //完成时间
	ClientName        string `bson:"client_name"`              // 客户名称
	CreatedBy         uint   `bson:"created_by"`               // 创建人
	CreatedAt         string `bson:"created_at"`               // 创建时间
	ReportUrl         string `bson:"report_url"`               //报告地址
	CallbackUrl       string `bson:"callback_url"`             //回调地址
	SspTaskId         uint   `bson:"ssp_task_id"`              //task_id  int类型
	TaskTemplateName  string `bson:"task_template_name"`       //任务模板名字
	WorkUID           string `bson:"work_uid"`                 // work uid
	FileSh256         string `bson:"file_sh256"`               //上传文件的sh256
	FileMd5           string `json:"file_md5" bson:"file_md5"` //上传文件的md5
	PkgName           string `bson:"pkg_name"`
	CertMd5           string `bson:"cert_md_5"`
	Datasource        int    `json:"datasource"`
	Edited            int    `json:"edited"`
	EngineIp          string `bson:"engine_ip"`
	FinishedItemCount int    `bson:"finished_item_count"`
	IsDeleted         int    `bson:"is_deleted"`
	IsVersionTarget   int    `bson:"is_version_target"`
	//CreatedAt string `json:"created_at"`
	UpdatedAt string `bson:"updated_at"`
	//ItemCount int `json:"item_count"`
	TaskUid    string `bson:"task_uid"`
	Key        string `json:"key"`
	State      string `json:"state"`
	Md5        string `bson:"md5"`
	ProductKey string `bson:"product_key"`
}

func Srctask() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var ceping_src_task1 []ceping_src_task
	err = db.Table("ceping_src_task").Find(&ceping_src_task1).Error
	if err != nil {
		fmt.Println(err)
	}

	db2	, err := gorm.Open("mysql", utils.Sspmysql())
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


	db3, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db3.Close() //用完之后关闭数据库连接

	for _, i2 := range ceping_src_task1 {
		var ce CPTaskMeta
		ce.TaskTemplateUID = "CP_TEMPLATE_SOURCE"
		ce.CallbackUrl = i2.CallbackUrl
		ce.IsDeleted = int(i2.IsDeleted[0])
		ce.CreatedAt = i2.TCommit
		ce.UpdatedAt = i2.TUpdate
		r := rand.New(rand.NewSource(time.Now().Unix()))
		//	fmt.Println(r.Int63())
		ce.SspTaskId = uint(r.Int63())
		ce.RequestID = i2.RequestId
		ce.Version = i2.SrcVersion
		ce.Md5 = i2.SrcMd5
		ce.FileMd5 = i2.SrcMd5
		ce.FileSize = int64(i2.SrcSize)
		ce.Type = "Source"
		ce.ProductKey = "SourceAudit"
		ce.FinishedTime = i2.TUpdate
		ce.FileName = i2.SrcFilename
		u1 := uuid.Must(uuid.NewV4(), nil)
		result := u1.String()
		result = strings.ReplaceAll(result, "-", "")
		ce.TaskUid = result
		ce.Key = "测评"
		if i2.FailMsg == "" {
			ce.State = "已完成"
		} else {
			ce.State = "失败"
		}

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

		err = db3.Table("ceping_user").Where("id = ?", i2.UserId).Last(&cepinguser).Error
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

		var aps_user struct{
			Id uint `json:"id"`
			Name string `json:"name"`
			RealName string `json:"real_name"`
			Password string `json:"password"`
		}
		err = db2.Table("aps_user").Where("name = ?", cepinguser.Username).Last(&aps_user).Error
		if err != nil {
			fmt.Println(err)
		}
		ce.CreatedBy = aps_user.Id

		//report_url
		//report_lang
		//score

		if i2.ResCode == 0 {
			ce.Result = "安全"
		}
		if i2.ResCode ==1 {
			ce.Result = "异常"
		}
		if i2.ResCode ==2 {
			ce.Result = "不安全"
		}


		c := session.DB("ssp").C("tws_task")
		err = c.Insert(&ce)
		if err != nil {
			log.Fatal(err)
		}

		//传入对应任务的  taskid（任务id，迁移的表）   taskuid（自己生成的）
		//把对应任务的回调数据存入mongodb
		fmt.Println(23443)
		//	Ipataskresult1(i2.Id, result, i2.RequestId)
		Srctaskresult(i2.Id, result, i2.RequestId, int(uint(r.Int63())))
		break
	}

}
