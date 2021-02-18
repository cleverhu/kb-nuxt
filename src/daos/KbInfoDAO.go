package daos

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/models/DocModel"
	"strings"
)

type KbInfoDAO struct {
	DB  *gorm.DB      `inject:"-"`
	Rds *redis.Client `inject:"-"`
}

func NewKbInfoDao() *KbInfoDAO {
	return &KbInfoDAO{}
}

func (this *KbInfoDAO) GetKbDetail(username string, kbName string) string {
	str, _ := this.Rds.Get("kb:120").Result()
	return str
}


func (this *KbInfoDAO) GetDocDetail(shortUrl string) *DocModel.DocContent {
	result := &DocModel.DocContent{}
	if strings.HasPrefix(shortUrl,"_grp"){
		this.DB.Table("docs").Raw("select docs.doc_title as doc_title,docs.doc_content as doc_content from doc_grps LEFT JOIN  docs on doc_grps.group_id  = docs.group_id WHERE doc_grps.shorturl = ?", shortUrl[4:]).First(&result)
	}else 	if strings.HasPrefix(shortUrl,"_kb"){
		this.DB.Table("kbs").Raw("select kb_name as doc_title,kb_desc as doc_content from kbs where kb_id = ?", shortUrl[3:]).First(&result)
	} else{
		this.DB.Table("docs").Raw("select doc_title,doc_content from docs where shorturl = ?", shortUrl).Find(&result)
	}

	return result
}
