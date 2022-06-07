package model

import (
	"github.com/amingze/gochive/db"
	"time"
)

var tableName = "memo"

type MemoModel struct {
	db db.DB
}

type Memo struct {
	Id         string
	CreateTime time.Time
	UpdateTime time.Time
	Content    string
}

func (mm MemoModel) Add(memo Memo) error {
	return mm.db.Add(tableName, memo.Id, &memo)
}

func (mm MemoModel) Read(memo Memo) error {
	return mm.db.Read(tableName, memo.Id, &memo)
}
