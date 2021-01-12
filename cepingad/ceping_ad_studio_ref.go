package cepingad

import (
	"awesomeProject6/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ceping_ad_studio_ref struct {
	Id int `json:"id" gorm:"primary_key"`
	CategoryKey string `json:"category_key" gorm:"primary_key"`
	ItemKey string `json:"item_key" gorm:"primary_key"`
	Sort int `json:"sort"`
	CategorySort int `json:"category_sort"`
}

func Studioref()  {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var ceping_ad_studio_ref1 []ceping_ad_studio_ref
	err = db.Table("ceping_ad_studio_ref").Find(&ceping_ad_studio_ref1).Error
	if err != nil {
		fmt.Println(err)
	}


	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range ceping_ad_studio_ref1 {
		fmt.Println(i2.CategoryKey + "" + i2.ItemKey)
		var ceping_ad_studio_ref2 ceping_ad_studio_ref
		ceping_ad_studio_ref2.Id = i2.Id
		ceping_ad_studio_ref2.ItemKey = i2.ItemKey
		ceping_ad_studio_ref2.Sort = i2.Sort
		ceping_ad_studio_ref2.CategorySort = i2.CategorySort
		ceping_ad_studio_ref2.CategoryKey = i2.CategoryKey
		err = db2.Table("ceping_ad_studio_ref").Create(&ceping_ad_studio_ref2).Error
		if err != nil {
			fmt.Println(err)
		}
		break
	}
}
