package configuration

import (
	"knowledgeBaseNuxt/src/daos"
	"knowledgeBaseNuxt/src/services"
)

type KbInfoServiceConfig struct {
}

func NewKbInfoServiceConfig() *KbInfoServiceConfig {
	return &KbInfoServiceConfig{}
}

func (this *KbInfoServiceConfig) KbUserDAO() *daos.KbInfoDAO {
	return daos.NewKbInfoDao()
}

func (this *KbInfoServiceConfig) KbUserService() *services.KbInfoService {
	return services.NewKbInfoService()
}
