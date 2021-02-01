package main

import (
	// 初始化路由
	_ "github.com/handsomestWei/learn-beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func main() {
	setSwaggerUI()
	setDefaultLog()
	setFilters()
	// 启动
	beego.Run()
}

// 配置swagger UI
func setSwaggerUI() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

// 配置日志
func setDefaultLog() {
	// 设置日志级别
	logs.SetLevel(logs.LevelDebug)
	// 设置日志输出格式：打印行号和方法名
	logs.SetLogFuncCall(true)
}

// 配置过滤器
func setFilters() {
	beego.InsertFilter("/api/v1/cms/user", beego.BeforeRouter, func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/")
		}
	})
}

