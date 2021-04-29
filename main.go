package main

import (
	"Shield/model"
	"Shield/routes"
)

func main() {
	// 引用数据库
	model.Init()
	// 引入路由组件
	routes.InitRouter()

}
