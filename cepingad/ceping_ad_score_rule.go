package cepingad

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ceping_ad_score_rule struct {
	ItemKey string `json:"item_key"`
	Max int `json:"max"`
	Min int `json:"min"`
	ResultType int `json:"result_type"`
	Score int `json:"score"`
	Val string `json:"val"`
}

func Scorerule()  {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var ceping_ad_score_rule1 []ceping_ad_score_rule
	err = db.Table("ceping_ad_score_rule").Find(&ceping_ad_score_rule1).Error
	if err != nil {
		fmt.Println(err)
	}


	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range ceping_ad_score_rule1 {
		var ceping_ad_score_rule2 ceping_ad_score_rule
	    ceping_ad_score_rule2.ItemKey = i2.ItemKey
	    ceping_ad_score_rule2.Max = i2.Max
	    ceping_ad_score_rule2.Min = i2.Min
	    ceping_ad_score_rule2.ResultType = i2.ResultType
	    ceping_ad_score_rule2.Score = i2.Score
	    ceping_ad_score_rule2.Val = i2.Val
		err := db2.Table("ceping_ad_score_rule").Create(&ceping_ad_score_rule2).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}
