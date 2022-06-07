package db

import (
	"encoding/json"
	"errors"
	"github.com/amingze/gochive/utils/fileutil"
	"path"
)

type DB struct {
	Path string
}

func New(fileName string) (*DB, error) {
	db := &DB{}
	_, err := fileutil.MkFile(fileName)
	db.Path = fileName
	return db, err
}

func (db DB) Add(table string, key string, model any) error {
	file, err := fileutil.MkFile(path.Join(table, key))
	if err != nil {
		return err
	}
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}
func (db DB) Read(table string, key string, model any) (err error) {
	exist := fileutil.Exist(path.Join(table, key))
	if exist {
		b, err := fileutil.Open(path.Join(table, key))
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &model)
	} else {
		err = errors.New("读取的文件不存在")
	}
	return
}

func (db DB) Exist(table string, key string, model any) {

}

func (db DB) Delete(table string, key string) {

}
func (db DB) ReadAll(table string, key string) {

}
