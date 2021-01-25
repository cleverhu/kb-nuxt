package DocModel

type DocImpl struct {
	DocID       int           `gorm:"column:doc_id;primary_key"  json:"doc_id"`
	DocTitle    string        `gorm:"column:doc_title" json:"label"`
	DocShortUrl string        `gorm:"column:shorturl" json:"url"`
	DocHref     string        `gorm:"-"` //拼接地址字符串
	Children    []interface{} `gorm:"-" json:"children,omitempty"`
}

type DocGrpImpl struct {
	GroupID       int           `gorm:"column:group_id;primary_key" json:"group_id"`
	GroupName     string        `gorm:"column:group_name" json:"label"`
	GroupShortUrl string        `gorm:"column:shorturl" json:"url"`
	Children      []interface{} `gorm:"-" json:"children,omitempty"`
}
