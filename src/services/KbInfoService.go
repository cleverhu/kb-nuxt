package services

import (
	"knowledgeBaseNuxt/src/daos"
	"knowledgeBaseNuxt/src/models/DocModel"
)

type KbInfoService struct {
	KbUserDao *daos.KbInfoDAO `inject:"-"`
}

func NewKbInfoService() *KbInfoService {
	return &KbInfoService{}
}

func (this *KbInfoService) KbDetailByID(username, kbName string) string {
	return this.KbUserDao.GetKbDetail(username, kbName)
}

func (this *KbInfoService) DocDetailByID(shortUrl string) *DocModel.DocContent {
	return this.KbUserDao.GetDocDetail(shortUrl)
}
