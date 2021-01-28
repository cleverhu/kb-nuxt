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
	db            *gorm.DB                `inject:"-"`
	KbInfoService *services.KbInfoService `inject:"-"`
	rds           *redis.Client           `inject:"-"`
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
	result, err := this.rds.Get("kb:" + id).Result()
	goft.Error(err, "获取知识库失败")
	return gin.H{"code": 10000, "result": result}
}

func (this *KbInfoController) Build(goft *goft.Goft) {
	goft.HandleWithFairing("GET", "/kns/:id", this.KbDetailByID)

}
