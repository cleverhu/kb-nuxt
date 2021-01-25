package DocGrpModel

import "knowledgeBaseNuxt/src/models/DocModel"

type DocGrpImpl struct {
	GroupID       int                 `gorm:"column:group_id;primary_key" json:"-"`
	GroupName     string              `gorm:"column:group_name" json:"label"`
	GroupShortUrl string              `gorm:"column:shorturl" json:"url"`
	Children      []*DocModel.DocImpl `gorm:"-" json:"children,omitempty"`
	//zDocs          []*DocModel.DocImpl `gorm:"-" json:"children_docs,omitempty"`
}
//[]*DocGrpImpl       `gorm:"-" json:"children,omitempty"`

func New(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {
	d := &DocGrpImpl{}
	DocGrpModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocGrpImpl) Mutate(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {

	DocGrpModelAttrFuncs(attrs).Apply(this)
	return this
}
