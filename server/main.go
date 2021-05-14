package main

import (
	"gitee.com/favefan/doconsole/core"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/initialize"
	_ "gitee.com/favefan/doconsole/models"
)

func main() {
	initialize.SettingSetup()
	global.GDB = initialize.DBSetup()
	//global.GClient = initialize.DockerClientSetup()
	if global.GDB != nil {
		initialize.TableDBInit(global.GDB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GDB.DB()
		defer db.Close()
	}

	core.RunServer()
}
