package daos

import (
	"fmt"
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/models/DocModel"
)

type KbInfoDAO struct {
	DB *gorm.DB `inject:"-"`
}

func NewKbInfoDao() *KbInfoDAO {
	return &KbInfoDAO{}
}

func (this *KbInfoDAO) GetKbDetail(username string, kbName string) []*DocModel.DocGrpImpl {
	var dgm []*DocModel.DocGrpImpl

	kbID := 120

	this.getKbDetail("test", kbID, 0, &dgm)
	return dgm
}

func (this *KbInfoDAO) getKbDetail(kbName string, kbID, groupID int, result *[]*DocModel.DocGrpImpl) []*DocModel.DocGrpImpl {
	var temp  []*DocModel.DocGrpImpl
	this.DB.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order `, kbID, groupID).Find(&temp)
	for _, v := range temp {
		*result=append(*result,v)
		var grp2  []*DocModel.DocGrpImpl
		this.getKbDetail(kbName, kbID, v.GroupID,&grp2)

		//for _, item := range grp2 {
		//	*result=append(*result,item)
		//}
		fmt.Println(*result)
		var docs []*DocModel.DocImpl
		this.DB.Table("docs").Raw(`select doc_id,doc_title,shorturl from docs 
where kb_id = ? and group_id = ? 
order by  doc_id`, kbID, v.GroupID).Find(&docs)
		for _, doc := range docs {


			doc.DocHref = "/" + kbName + "/" + v.GroupShortUrl + "/" + doc.DocShortUrl
			v.Children = append(v.Children,doc)
			var grp  []*DocModel.DocGrpImpl
			this.getKbDetail(kbName, kbID, v.GroupID,&grp)

			for _, item := range grp {
				v.Children =append(v.Children ,item)
			}
		}



	}
		fmt.Println(result)
	return *result
}
