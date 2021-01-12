package cepingsrc

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ceping_src_audit_item struct {
	AuditCategory string `json:"audit_category"`
	AuditDef      string `json:"audit_def"`
	AuditDesc     string `json:"audit_desc"`
	AuditName     string `json:"audit_name"`
	CodeExample   string `json:"code_example"`
	Level         string `json:"level"`
	RuleKey       string `json:"rule_key"`
	Score         int    `json:"score"`
	See           string `json:"see"`
	Solution      string `json:"solution"`
	LangType      string `json:"lang_type"`
}

type ceping_src_audit_item0 struct {
	AuditCategory string `json:"audit_category"`
	AuditDef      string `json:"audit_def"`
	AuditDesc     string `json:"audit_desc"`
	AuditName     string `json:"audit_name"`
	CodeExample   string `json:"code_example"`
	Level         string `json:"level"`
	RuleKey       string `json:"rule_key"`
	Score         int    `json:"score"`
	See           string `json:"see"`
	Solution      string `json:"solution"`
	LangType      string `json:"lang_type"`
}

func Srcaudititem() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var cepingsrcaudititem []ceping_src_audit_item
	err = db.Table("ceping_src_audit_item").Find(&cepingsrcaudititem).Error
	if err != nil {
		fmt.Println(err)
	}

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range cepingsrcaudititem {
		ceping_src_audit_item2 := ceping_src_audit_item0{}
		ceping_src_audit_item2.AuditCategory = i2.AuditCategory
		ceping_src_audit_item2.AuditDef = i2.AuditDef
		ceping_src_audit_item2.AuditDesc = i2.AuditDesc
		ceping_src_audit_item2.AuditName = i2.AuditName
		ceping_src_audit_item2.CodeExample = i2.CodeExample
		ceping_src_audit_item2.Level = i2.Level
		ceping_src_audit_item2.RuleKey = i2.RuleKey
		ceping_src_audit_item2.Score = i2.Score
		ceping_src_audit_item2.See = i2.See
		ceping_src_audit_item2.Solution = i2.Solution
		ceping_src_audit_item2.LangType = i2.LangType

		//ceping_ipa_audit_item2.AdminSetting = 0
		err = db2.Table("ceping_src_audit_item").Create(&ceping_src_audit_item2).Error
		if err != nil {
			fmt.Println(err)
		}
		break
	}

}
