package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/services"
	"strconv"
)

type KbInfoController struct {
	DB            *gorm.DB                `inject:"-"`
	KbInfoService *services.KbInfoService `inject:"-"`
	Rds           *redis.Client           `inject:"-"`
}

func NewKbInfoController() *KbInfoController {
	return &KbInfoController{}
}

func (this *KbInfoController) Name() string {
	return "KbUserController"
}

func (this *KbInfoController) KbDetailByID(ctx *gin.Context) goft.Json {
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	goft.Error(err, "知识库ID错误")
	return gin.H{"code": 10000, "result": this.KbInfoService.KbDetailByID("", "")}
}

func (this *KbInfoController) DocDetailByID(ctx *gin.Context) goft.Json {
	shortUrl := ctx.Param("shortUrl")
	return gin.H{"code": 10000, "result": this.KbInfoService.DocDetailByID(shortUrl)}
}

func (this *KbInfoController) Build(goft *goft.Goft) {
	//获取知识库分组信息
	goft.Handle("GET", "/kns/:id", this.KbDetailByID).
		//获取文章
		Handle("GET", "/doc/:shortUrl", this.DocDetailByID)

}
