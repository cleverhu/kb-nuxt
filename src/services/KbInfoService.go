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

func (this *KbInfoService) KbDetailByID(kbID int) []*DocModel.DocGrpImpl {
	return this.KbUserDao.GetKbDetail("", "test")
}
