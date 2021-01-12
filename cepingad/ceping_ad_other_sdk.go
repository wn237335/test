package cepingad

import (
	"awesomeProject6/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"time"
)

type ceping_ad_other_sdk struct {
	Id      int    `json:"id"`
	SdkName string `json:"sdk_name"`
	SdkPkgs string `json:"sdk_pkgs"`
	Enable  []byte `json:"enable"`
	Type    string `json:"type"`
}

type Othersdk struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	SdkName   string `json:"sdk_name"`
	SdkPkgs   string `json:"sdk_pkgs"`
	Enable    bool   `json:"enable"`
	Type      string `json:"type"`
}

type Othersdk1 struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	//DeletedAt string `json:"deleted_at"`
	SdkName string `json:"sdk_name"`
	SdkPkgs string `json:"sdk_pkgs"`
	Enabled int    `json:"enabled"`
	Type    string `json:"type"`
}

var ctx = context.Background()

func Cepingothersdk() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	var ceping_ad_other_sdk1 []ceping_ad_other_sdk
	err = db.Table("ceping_ad_other_sdk").Find(&ceping_ad_other_sdk1).Error
	if err != nil {
		fmt.Println(err)
	}

	for _, i2 := range ceping_ad_other_sdk1 {
		var oth Othersdk1
		oth.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		oth.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		oth.SdkPkgs = i2.SdkPkgs
		oth.SdkName = i2.SdkName
		oth.Enabled = int(i2.Enable[0])
		oth.Type = i2.Type
		err := db2.Table("other_sdk").Create(&oth).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	var othersdka []Othersdk1
	err = db2.Table("other_sdk").Find(&othersdka).Error
	if err != nil {
		fmt.Println(err)
	}

	var items []Othersdk
	for _, i2 := range othersdka {
		var othersdk Othersdk
		othersdk.Id = i2.Id
		othersdk.CreatedAt = time.Now().Format(time.RFC3339)
		othersdk.SdkName = i2.SdkName
		othersdk.SdkPkgs = i2.SdkPkgs

		if i2.Enabled == 1 {
			othersdk.Enable = true
		}
		othersdk.Type = i2.Type
		items = append(items, othersdk)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.Sspredisserverandport(),
		Password: utils.Sspredispwd(), // no password set
		DB:       0,                   // use default DB
	})

	/*val, err := rdb.Get(ctx, "ceping_othersdk").Result()
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(val)
	val1 := make([]Othersdk, 0)
	json.Unmarshal([]byte(val), &val1)
	for _, i2 := range val1 {
		items = append(items, i2)
	}*/

	ss, _ := json.Marshal(items)

	err = rdb.Set(ctx, "ceping_othersdk", ss, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}
