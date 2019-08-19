package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"web/controllers"
)

func main(){
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册控制器
	mvc.New(app.Party("hello")).Handle(new(controllers.MovieController))
	//注册模板目录，views目录下所有html文件都当成模板
	app.RegisterView(iris.HTML("./src/web/views",".html"))
	app.Run(iris.Addr("localhost:8080"))
}
