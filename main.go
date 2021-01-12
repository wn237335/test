package main

import "awesomeProject6/cepingad"

func main() {

	//1、用户数据迁移
	//cepinguser.Cpuser()

	//2.cepingad数据迁移
	//cepingad.Audititem()
	cepingad.Cepingothersdk()
	//cepingad.Scorerule()
	//cepingad.Sdkresult()
	//cepingad.Settingitem1()
	//cepingad.Studioref()
	//cepingad.Adtask()   //cepingad.Taskresult()

	//3.cepingipa数据迁移
	//cepingipa.Ipaaudititem()
	//cepingipa.Ipasettingitem1()
	//cepingipa.Ipatask() //在迁移每条ios的任务时，同时迁移回调数据     //cepingipa.Ipataskresult()

	//4.cepingmini
	//cepingmini.Miniprogramsettingitem1()
	//cepingmini.Miniprogramsettingitem()

	//5.cepingnet
	//cepingnet.Nettask()

	//6.cepingsdk
	//cepingsdk.Sdkaudititem()
	//cepingsdk.Sdksettingitem1()
	//cepingsdk.Sdktask()

	//7.cepingsrc
	//cepingsrc.Srcaudititem()
	//cepingsrc.Srctask()

	//8.cepingother
	//cepingother.Whitelist()

}
