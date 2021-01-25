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
	//先找到分组
	this.DB.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order`, kbID, groupID).Find(&result)

	//遍历分组找文档
	for _, v := range *result {
		//找到文档
		var docs []*DocModel.DocImpl
		this.DB.Table("docs").Raw(`select doc_id,doc_title,shorturl from docs 
where kb_id = ? and group_id = ? 
order by doc_id`, kbID, v.GroupID).Find(&docs)

		for _, doc := range docs {
			//给文档追加到子元素
			doc.DocHref = "/" + kbName + "/" + v.GroupShortUrl + "/" + doc.DocShortUrl
			v.Children = append(v.Children, doc)
		}

		//寻找子分组 这个数据是临时的，不会返回真实数据
		var grp []*DocGrpModel.DocGrpImpl
		this.getKbDetail(kbName, kbID, v.GroupID, &grp)
		for _, item := range grp {
			v.Children = append(v.Children, item)
		}
	}
	return *result
}


