package model

import (
	"github.com/amingze/gochive/db"
	"testing"
	"time"
)

func TestMemoModel_Add(t *testing.T) {
	type fields struct {
		db db.DB
	}
	type args struct {
		memo Memo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "添加测试",
			fields: fields{
				db.DB{Path: "./test/hi"},
			},
			args: args{
				memo: Memo{
					Id:         "1",
					Content:    "Hi~ o(*￣▽￣*)ブ你好",
					CreateTime: time.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := MemoModel{
				db: tt.fields.db,
			}
			mm.Add(tt.args.memo)
		})
	}
}
