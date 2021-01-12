package cepinguser

import (
	"awesomeProject6/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"strings"
	"time"
)

/*
测评4.6.0.1  用户数据迁移
ceping_user    =>   admin_user   client   client_template_mapper   client_user
*/

type cepinguser struct {
	AndroidStudioId   int    `json:"android_studio_id"`
	ApiSignature      string `json:"api_signature"`
	AuthType          string `json:"auth_type"`
	CepingApps        int    `json:"ceping_apps"`
	CepingTimes       int    `json:"ceping_times"`
	Company           string `json:"company"`
	ContactEmail      string `json:"contact_email"`
	Enable            []byte `json:"enable"`
	IosEndtime        string `json:"ios_endtime"`
	IosStudioId       int    `json:"ios_studio_id"`
	IpaApps           int    `json:"ipa_apps"`
	IpaTimes          int    `json:"ipa_times"`
	IsFirstlogin      []byte `json:"is_firstlogin"`
	LastUseday        int    `json:"last_useday"`
	NetEndtime        string `json:"net_endtime"`
	NetTimes          int    `json:"net_times"`
	Password          string `json:"password"`
	Platform          string `json:"platform"`
	Platforms         string `json:"platforms"`
	ReportFormats     string `json:"report_formats"`
	ReportIosLanguage int    `json:"report_ios_language"`
	ReportLanguage    int    `json:"report_language"`
	//	ReportLanguages           int    `json:"report_languages"`
	ReportTemplate            string `json:"report_template"`
	RoleType                  string `json:"role_type"`
	SdkApps                   int    `json:"sdk_apps"`
	SdkEndtime                string `json:"sdk_endtime"`
	SdkTimes                  int    `json:"sdk_times"`
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

type Prou struct {
	SystemID     uint     `json:"system_id"`         //关联的系统id
	Name         string   `json:"name"`              //授权的服务（产品）名
	Key          string   `json:"key"`               //授权的服务产品key
	UsageLimit   int64    `json:"usage_limit"`       //授权产品的使用次数
	UsageCounter int64    `json:"usage_counter"`     //产品使用的次数
	SampleLimit  int64    `json:"sample_limit"`      // 样本限制
	SampleUsed   int64    `json:"sample_used"`       // 使用样本数量
	Features     []string `json:"features" gorm:"-"` //产品授权的功能特点数组
	StartedAt    string   `json:"started_at"`        // 授权开始时间
	EndedAt      string   `json:"ended_at"`          // 授权结束时间
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func Cpuser() {

	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var cepinguser []cepinguser
	err = db.Table("ceping_user").Find(&cepinguser).Error
	if err != nil {
		fmt.Println(err)
	}

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	var admin_user struct {
		Name     string `json:"name"`
		RealName string `json:"real_name"`
		Password string `json:"password"`
		Locked   int    `json:"locked"`
		Fails    int    `json:"fails"`
	}

	var client struct {
		CreatedAt        string `json:"created_at"`
		UpdatedAt        string `json:"upadted_at"`
		Name             string `json:"name"`
		ClientKey        string `json:"client_key"`
		AdminAccount     string `json:"admin_account"`
		Expiration       string `json:"expiration"`
		Status           int    `json:"status"`
		SampleCountLimit int    `json:"sample_count_limit"`
		SampleSizeLimit  int    `json:"sample_size_limit"`
	}

	var client_template_mapper struct {
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"upadted_at"`
		ClientKey    string `json:"client_key"`
		ClientName   string `json:"client_name"`
		TemplateId   int    `json:"template_id"`
		TemplateName string `json:"template_name"`
	}

	var client_user struct {
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"upadted_at"`
		ClientKey    string `json:"client_key"`
		ClientName   string `json:"client_name"`
		UserAccount  string `json:"user_account"`
		UserRealName string `json:"user_real_name"`
		Status       int    `json:"status"`
		RoleId       int    `json:"role_id"`
		RoleName     string `json:"role_name"`
		Email        string `json:"email"`
		IsAdmin      bool   `json:"is_admin"`
		Products     string `json:"products"`
		Products1    []struct {
			SystemID     uint     `json:"system_id"`         //关联的系统id
			Name         string   `json:"name"`              //授权的服务（产品）名
			Key          string   `json:"key"`               //授权的服务产品key
			UsageLimit   int64    `json:"usage_limit"`       //授权产品的使用次数
			UsageCounter int64    `json:"usage_counter"`     //产品使用的次数
			SampleLimit  int64    `json:"sample_limit"`      // 样本限制
			SampleUsed   int64    `json:"sample_used"`       // 使用样本数量
			Features     []string `json:"features" gorm:"-"` //产品授权的功能特点数组
			StartedAt    string   `json:"started_at"`        // 授权开始时间
			EndedAt      string   `json:"ended_at"`          // 授权结束时间
		} `json:"products1"`
	}

	var aps_user struct {
		Name     string `json:"name"`
		RealName string `json:"real_name"`
		Password string `json:"password"`
	}

	//fmt.Println(string(b))
	for _, i2 := range cepinguser {
		b := make([]rune, 32)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		key := fmt.Sprintf("ct-%s", string(b))
		//admin_user   client   client_template_mapper   client_user
		admin_user.Name = i2.Username
		admin_user.RealName = i2.Username
		admin_user.Password = i2.Password
		admin_user.Locked = 0
		admin_user.Fails = 0
		err = db2.Table("admin_user").Create(&admin_user).Error
		if err != nil {
			fmt.Println(err)
		}

		client.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		client.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		client.Name = i2.Username
		client.ClientKey = key
		client.AdminAccount = i2.Username
		client.Expiration = "" //账户过期时间
		client.Status = 1
		client.SampleCountLimit = 0
		client.SampleSizeLimit = 0
		err = db2.Table("client").Create(&client).Error
		if err != nil {
			fmt.Println(err)
		}

		client_template_mapper.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		client_template_mapper.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		client_template_mapper.ClientKey = key
		client_template_mapper.ClientName = i2.Username
		client_template_mapper.TemplateId = 1
		client_template_mapper.TemplateName = "默认"
		err = db2.Table("client_template_mapper").Create(&client_template_mapper).Error
		if err != nil {
			fmt.Println(err)
		}

		client_user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		client_user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		client_user.ClientKey = key
		client_user.ClientName = i2.Username
		client_user.UserAccount = i2.Username
		client_user.UserRealName = i2.Username
		client_user.Status = 1
		client_user.RoleId = 0
		client_user.RoleName = ""
		client_user.Email = i2.ContactEmail
		client_user.IsAdmin = true

		spla := strings.Split(i2.Platforms, ",")

		i := 1
		for {
			if len(spla)-i < 0 {
				break
			}
			sspl := spla[len(spla)-i]
			if sspl == "android" {
				if i2.TEndtime != "" {
					var prou1 Prou
					prou1.Name = "Android测评"
					prou1.Key = "AndroidAudit"
					prou1.EndedAt = i2.TEndtime
					prou1.SystemID = 1
					prou1.SampleLimit = int64(i2.CepingApps)
					prou1.UsageCounter = int64(i2.CepingTimes)
					client_user.Products1 = append(client_user.Products1, prou1)
				}
			}
			if sspl == "ios" {
				if i2.IosEndtime != "" {
					var prou2 Prou
					prou2.Name = "iOS测评"
					prou2.Key = "IOSAudit"
					prou2.EndedAt = i2.IosEndtime
					prou2.SystemID = 1
					prou2.SampleLimit = int64(i2.IpaApps)
					prou2.UsageCounter = int64(i2.IpaTimes)
					client_user.Products1 = append(client_user.Products1, prou2)
				}
			}
			if sspl == "net" {
				if i2.NetEndtime != "" {
					var prou3 Prou
					prou3.Name = "web测评"
					prou3.Key = "WebAudit"
					prou3.EndedAt = i2.NetEndtime
					prou3.SystemID = 1
					prou3.UsageCounter = int64(i2.NetTimes)
					client_user.Products1 = append(client_user.Products1, prou3)
				}
			}
			if sspl == "sdk" {
				if i2.SdkEndtime != "" {
					var prou4 Prou
					prou4.Name = "SDK测评"
					prou4.Key = "SDKAudit"
					prou4.EndedAt = i2.SdkEndtime
					prou4.SystemID = 1
					prou4.SampleLimit = int64(i2.SdkApps)
					prou4.UsageCounter = int64(i2.SdkTimes)
					client_user.Products1 = append(client_user.Products1, prou4)
				}
			}
			if sspl == "source" {
				if i2.SourceEndTime != "" {
					var prou5 Prou
					prou5.Name = "源码测评"
					prou5.Key = "SourceAudit"
					prou5.EndedAt = i2.SourceEndTime
					prou5.SystemID = 1
					prou5.UsageCounter = int64(i2.SourceTimes)
					client_user.Products1 = append(client_user.Products1, prou5)
				}
			}
			if sspl == "miniprogram" {
				if i2.MiniProgramEndtime != "" {
					var prou6 Prou
					prou6.Name = "小程序测评"
					prou6.Key = "MiniProgramAudit"
					prou6.EndedAt = i2.MiniProgramEndtime
					prou6.SystemID = 1
					prou6.SampleLimit = int64(i2.MiniProgramApps)
					prou6.UsageCounter = int64(i2.MiniProgramTimes)
					client_user.Products1 = append(client_user.Products1, prou6)
				}
			}
			i++
		}

		//android,ios,source,net,sdk,miniprogram
		fmt.Println(client_user.Products)
		ssaa, _ := json.Marshal(client_user.Products1)
		client_user.Products = string(ssaa)
		err = db2.Table("client_user").Create(&client_user).Error
		if err != nil {
			fmt.Println(err)
		}

		aps_user.Name = i2.Username
		aps_user.Password = i2.Password
		aps_user.RealName = i2.Username
		err = db2.Table("aps_user").Create(&aps_user).Error
		if err != nil {
			fmt.Println(err)
		}
		break
	}
}
