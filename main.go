package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	_ "knowledgeBaseNuxt/src/common"
	"knowledgeBaseNuxt/src/configuration"
	"knowledgeBaseNuxt/src/controllers"
)

func main() {
	//fmt.Println("123")
	goft.Ignite().
		Config(configuration.NewDBConfig(), configuration.NewKbInfoServiceConfig(),configuration.NewRedisConfig()).
		Attach().//全局中间件
		Mount("", controllers.NewKbInfoController()).
		Launch()
}
