package model

import "gorm.io/gorm"

type Memo struct {
	Model
	Content string    `json:"content"`
	Uid     int64     `json:"uid" gorm:"unique_index;not null"`
	Matter  []*Matter `json:"matter,omitempty" gorm:"many2many:memo_matters;"`
	Tags    []*Tag    `json:"tags" gorm:"many2many:memo_tags;"`
}
type MemoTag struct {
	gorm.Model
}

func (Memo) TableName() string {
	return "memo"
}
func NewMemo() *Memo {
	return &Memo{}
}
