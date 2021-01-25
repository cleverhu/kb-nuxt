package DocModel

import "fmt"

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

func (this *DocImpl) String() string {
	return fmt.Sprintf("DocID:%d,DocTitle:%s,DocShortUrl:%s,DocHref:%s", this.DocID, this.DocTitle, this.DocShortUrl, this.DocHref)
}

func (this *DocGrpImpl) String() string {
	return fmt.Sprintf("GroupID:%d,GroupName:%s,GroupShortUrl:%s", this.GroupID, this.GroupName, this.GroupShortUrl)
}
