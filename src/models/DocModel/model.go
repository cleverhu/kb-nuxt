package DocModel

type DocImpl struct {
	DocID       int    `gorm:"column:doc_id;primary_key"`
	DocTitle    string `gorm:"column:doc_title"`
	DocShortUrl string `gorm:"column:shorturl"`
	DocHref     string `gorm:"-"` //拼接地址字符串
	//Children []*DocGrpModel.DocGrpImpl       `gorm:"-" json:"children,omitempty"`
}

func New(attrs ...DocModelAttrFunc) *DocImpl {
	d := &DocImpl{}
	DocModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocImpl) Mutate(attrs ...DocModelAttrFunc) *DocImpl {
	DocModelAttrFuncs(attrs).Apply(this)
	return this
}
