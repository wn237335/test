package cepingad

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ceping_ad_sdk_result struct {
	ItemKey string `json:"item_key"`
	ItemValue string `json:"item_value"`
	SdkPoints int `json:"sdk_points"`
	SdkRiskNums int `json:"sdk_risk_nums"`
	Sdks string `json:"sdks"`
	TaskId int `json:"task_id"`
}

type ceping_ad_sdk_result0 struct {
	ItemKey string `json:"item_key"`
	ItemValue string `json:"item_value"`
	SdkPoints int `json:"sdk_points"`
	SdkRiskNums int `json:"sdk_risk_nums"`
	Sdks string `json:"sdks"`
	TaskUid string `json:"task_uid"`
}

func Sdkresult()  {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var ceping_ad_sdk_result1 []ceping_ad_sdk_result
	err = db.Table("ceping_ad_sdk_result").Find(&ceping_ad_sdk_result1).Error
	if err != nil {
		fmt.Println(err)
	}


	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range ceping_ad_sdk_result1 {
		var ceping_ad_sdk_result2 ceping_ad_sdk_result0
		ceping_ad_sdk_result2.ItemKey = i2.ItemKey
		ceping_ad_sdk_result2.ItemValue = i2.ItemValue
		ceping_ad_sdk_result2.SdkPoints = i2.SdkPoints
		ceping_ad_sdk_result2.SdkRiskNums = i2.SdkRiskNums
		ceping_ad_sdk_result2.Sdks = i2.Sdks
		ceping_ad_sdk_result2.TaskUid =strconv.Itoa(i2.TaskId)
		err := db2.Table("ceping_ad_sdk_result").Create(&ceping_ad_sdk_result2).Error
		if err != nil {
			fmt.Println(err)
		}
	}

}