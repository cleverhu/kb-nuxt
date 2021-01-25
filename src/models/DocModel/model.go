package DocModel

type DocImpl struct {
	DocID       int    `gorm:"column:doc_id;primary_key"`
	DocTitle    string `gorm:"column:doc_title"`
	DocShortUrl string `gorm:"column:shorturl"`
	DocHref     string `gorm:"-"` //拼接地址字符串
	Children []*DocGrpImpl       `gorm:"-" json:"children,omitempty"`
}

type DocGrpImpl struct {
	GroupID       int                 `gorm:"column:group_id;primary_key" json:"-"`
	GroupName     string              `gorm:"column:group_name" json:"label"`
	GroupShortUrl string              `gorm:"column:shorturl" json:"url"`
	Children      []*DocImpl `gorm:"-" json:"children,omitempty"`
}




