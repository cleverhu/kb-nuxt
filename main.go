package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	_ "knowledgeBaseNuxt/src/common"
	"knowledgeBaseNuxt/src/configuration"
	"knowledgeBaseNuxt/src/controllers"
)

func main() {
	goft.Ignite().
		Config(configuration.NewDBConfig(), configuration.NewKbInfoServiceConfig()).
		Attach().//全局中间件
		Mount("", controllers.NewKbInfoController()).
		Launch()
}
