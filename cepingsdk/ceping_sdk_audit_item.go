package cepingsdk

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ceping_sdk_audit_item struct {
	CategoryId int    `json:"category_id"`
	ItemKey    string `json:"item_key"`
	Level      string `json:"level"`
	Score      int    `json:"score"`
	Sort       int    `json:"sort"`
	Status     []byte `json:"status"`
	Solution   string `json:"solution"`
}

type ceping_sdk_audit_item0 struct {
	CategoryId int    `json:"category_id"`
	ItemKey    string `json:"item_key"`
	Level      string `json:"level"`
	Score      int    `json:"score"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	Solution   string `json:"solution"`
}

func  Sdkaudititem() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var cepingsdkaudititem []ceping_sdk_audit_item
	err = db.Table("ceping_sdk_audit_item").Find(&cepingsdkaudititem).Error
	if err != nil {
		fmt.Println(err)
	}

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range cepingsdkaudititem {
		ceping_sdk_audit_item2 := ceping_sdk_audit_item0{}
		ceping_sdk_audit_item2.CategoryId = i2.CategoryId
		ceping_sdk_audit_item2.ItemKey = i2.ItemKey
		ceping_sdk_audit_item2.Level = i2.Level
		ceping_sdk_audit_item2.Score = i2.Score
		ceping_sdk_audit_item2.Sort = i2.Sort
		ceping_sdk_audit_item2.Status = int(i2.Status[0])
		ceping_sdk_audit_item2.Solution = i2.Solution
		//ceping_ipa_audit_item2.AdminSetting = 0
		err = db2.Table("ceping_sdk_audit_item").Create(&ceping_sdk_audit_item2).Error
		if err != nil {
			fmt.Println(err)
		}
		break
	}

}

