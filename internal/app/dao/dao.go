package dao

import (
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/pkg/gormutil"
	"github.com/amingze/gochive/internal/pkg/utils/fileutil"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func Ready() bool {
	return gdb != nil
}

func Init(driver, dsn string) error {
	conf := gormutil.Config{
		Driver: driver,
		DSN:    dsn,
	}
	if driver == "" || driver == "sqlite" || driver == "sqlite3" {
		fileutil.MkFileAll(dsn)
	}
	db, err := gormutil.New(conf)
	if err != nil {
		return err
	}

	gdb = db.Debug()

	if err := gdb.AutoMigrate(model.Tables()...); err != nil {
		return err
	}

	return nil
}

func Transaction(fc func(tx *gorm.DB) error) error {
	return gdb.Transaction(fc)
}
