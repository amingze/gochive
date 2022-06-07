package fileutil

import (
	"io/ioutil"
	"os"
	"strings"
)

// MkDir 创建文件夹
func MkDir(path string) (err error) {
	_, statErr := os.Stat(path)
	if statErr == nil {
		return
	} else if os.IsNotExist(statErr) {
		err = os.MkdirAll(path, os.ModePerm)
		return
	} else {
		return statErr
	}
}

// MkFile 创建文件
func MkFile(path string) (file *os.File, err error) {
	splitPathFun := func(c rune) bool {
		if c == '\\' || c == '/' {
			return true
		} else {
			return false
		}
	}
	index := strings.LastIndexFunc(path, splitPathFun)
	dirs := path[:index]
	err = os.MkdirAll(dirs, os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err = os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Exist 判断文件存在
func Exist(path string) bool {
	_, statErr := os.Stat(path)
	if statErr == nil {
		return true
	} else {
		return false
	}
}

// NotExist 判断文件不存在
func NotExist(path string) bool {
	return !Exist(path)
}

// Open 打开文件
func Open(path string) (b []byte, err error) {
	b, err = ioutil.ReadFile(path)
	return
}
