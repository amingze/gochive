package model

type Tag struct {
	Model
	Name        string  `json:"name"`
	ParentTagId string  `json:"parent_tag_id"`
	Memos       []*Memo `json:"memos,omitempty" gorm:"many2many:memo_tags;"`
}

func (Tag) TableName() string {
	return "tag"
}
