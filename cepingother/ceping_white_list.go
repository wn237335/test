package cepingother

import (
	"awesomeProject6/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ceping_white_list struct {
	Id             int    `json:"id"`
	CreateTime     string `json:"create_time"`
	FileOrFunction string `json:"file_or_function"`
	ListType       string `json:"list_type"`
	Path           string `json:"path"`
	Type           string `json:"type"`
	Publish        []byte `json:"publish"`
}

type white_list struct {
	Id             int    `json:"id"`
	UpdatedAt      string `json:"updated_at"`
	CreatedAt      string `json:"created_at"`
	Type           string `json:"type"`
	ListType       string `json:"list_type"`
	Path           string `json:"path"`
	FileOrFunction string `json:"file_or_function"`
}

var ctx = context.Background()

func Whitelist() {
	db, err := gorm.Open("mysql", utils.Bangbangmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //用完之后关闭数据库连接
	//db.LogMode(true) //开启sql debug 模式
	var cepingwhitelist []ceping_white_list
	err = db.Table("ceping_white_list").Find(&cepingwhitelist).Error
	if err != nil {
		fmt.Println(err)
	}

	db2, err := gorm.Open("mysql", utils.Sspmysql())
	if err != nil {
		fmt.Println(err)
	}
	defer db2.Close() //用完之后关闭数据库连接

	for _, i2 := range cepingwhitelist {
		ceping_white_list2 := white_list{}
		ceping_white_list2.UpdatedAt = i2.CreateTime
		ceping_white_list2.CreatedAt = i2.CreateTime
		ceping_white_list2.FileOrFunction = i2.FileOrFunction
		ceping_white_list2.ListType = i2.ListType
		ceping_white_list2.Path = i2.ListType
		ceping_white_list2.Type = i2.Type
		err = db2.Table("white_list").Create(&ceping_white_list2).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	white_list2 := []white_list{}
	err = db2.Table("white_list").Find(&white_list2).Error
	if err != nil {
		fmt.Println(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.Sspredisserverandport(),
		Password: utils.Sspredispwd(), // no password set
		DB:       0,                   // use default DB
	})

	ss, _ := json.Marshal(white_list2)

	err = rdb.Set(ctx, "ceping_whitelist", ss, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}
