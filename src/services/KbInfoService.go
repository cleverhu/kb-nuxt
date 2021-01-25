package services

import (
	"knowledgeBaseNuxt/src/daos"
	"knowledgeBaseNuxt/src/models/DocGrpModel"
)

type KbInfoService struct {
	KbUserDao *daos.KbInfoDAO `inject:"-"`
}

func NewKbInfoService() *KbInfoService {
	return &KbInfoService{}
}

func (this *KbInfoService) KbDetailByID(kbID int) []*DocGrpModel.DocGrpImpl {
	return this.KbUserDao.GetKbDetail("", "test")
}
