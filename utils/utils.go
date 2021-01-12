package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

func Sspmysql() string {
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Sspmysql").Key("url").String())
	return iniConf.Section("Sspmysql").Key("url").String()
}

func Sspmongodbserver() string {
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Sspmongodb").Key("server").String())
	return iniConf.Section("Sspmongodb").Key("server").String()
}

func Sspmongodbname() string{
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Sspmongodb").Key("name").String())
	return iniConf.Section("Sspmongodb").Key("name").String()
}

func Sspmongodbuser() string{
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Sspmongodb").Key("user").String())
	return iniConf.Section("Sspmongodb").Key("user").String()
}

func Sspmongodbpwd() string{
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Sspmongodb").Key("pwd").String())
	return iniConf.Section("Sspmongodb").Key("pwd").String()
}


func Sspredisserverandport() string {
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s:%s",iniConf.Section("Sspredis").Key("server").String(),iniConf.Section("Sspredis").Key("port").String())
}

func Sspredispwd() string {
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	return iniConf.Section("Sspredis").Key("pwd").String()
}


func Bangbangmysql() string {
	iniConf, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iniConf.Section("Bangbangmysql").Key("url").String())
	return iniConf.Section("Bangbangmysql").Key("url").String()
}
