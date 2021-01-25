package daos

import (
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/models/DocGrpModel"
	"knowledgeBaseNuxt/src/models/DocModel"
)

type KbInfoDAO struct {
	DB *gorm.DB `inject:"-"`
}

func NewKbInfoDao() *KbInfoDAO {
	return &KbInfoDAO{}
}

func (this *KbInfoDAO) GetKbDetail(username string, kbName string) []*DocGrpModel.DocGrpImpl {
	var dgm []*DocGrpModel.DocGrpImpl

	kbID := 120

	this.getKbDetail("test", kbID, 0, &dgm)
	return dgm
}

func (this *KbInfoDAO) getKbDetail(kbName string, kbID, groupID int, result *[]*DocGrpModel.DocGrpImpl) []*DocGrpModel.DocGrpImpl {
	this.DB.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order `, kbID, groupID).Find(&result)
	for _, v := range *result {
		var docs []*DocModel.DocImpl
		this.DB.Table("docs").Raw(`select doc_id,doc_title,shorturl from docs 
where kb_id = ? and group_id = ? 
order by  doc_id`, kbID, groupID).Find(&docs)
		for _, doc := range docs {
			doc.DocHref = "/" + kbName + "/" + v.GroupShortUrl + "/" + doc.DocShortUrl
		}
		v.Docs = docs
		this.getKbDetail(kbName, kbID, v.GroupID, &v.Children)
	}
	return *result
}
