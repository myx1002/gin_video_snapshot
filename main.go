package main

import (
	"gin_video_snapshot/dao"
	"gin_video_snapshot/model"
	"gin_video_snapshot/routers"
)

func main() {
	// 创建数据库连接
	err := dao.InitMysql()
	if err != nil {
		panic(err.Error())
	}

	// 初始化model
	model.InitModel()

	// 初始化路由
	g := routers.InitRouter()
	_ = g.Run(":8100")
}
