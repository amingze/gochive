package fileutil

import (
	"os"
	"strings"
)

func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

//func Visit(dir string, suffix string, visitor func(filename string) error) error {
//	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//
//		if !info.IsDir() && strings.HasSuffix(path, suffix) {
//			return visitor(path)
//		}
//
//		return nil
//	})
//}

// MD5Hex returns the file md5 hash hex
//func MD5Hex(filepath string) string {
//	f, err := os.Open(filepath)
//	if err != nil {
//		return ""
//	}
//
//	md5hash := md5.New()
//	if _, err := io.Copy(md5hash, f); err != nil {
//		return ""
//	}
//	if err := f.Close(); err != nil {
//		logrus.Error(err)
//	}
//	return hex.EncodeToString(md5hash.Sum(nil)[:])
//}

// DetectContentType returns the file content-type
//func DetectContentType(filepath string) string {
//	mimeType := mime.TypeByExtension(path.Ext(filepath))
//	if mimeType != "" {
//		return mimeType
//	}
//
//	fileData, err := ioutil.ReadFile(filepath)
//	if err != nil {
//		return ""
//	}
//
//	return http.DetectContentType(fileData)
//}

//func UserHomeAbs(filename string) string {
//	if strings.HasPrefix(filename, "~/") {
//		u, _ := user.Current()
//		return filepath.Join(u.HomeDir, filename[1:])
//	}
//
//	return filename
//}

func MkFileAll(path string) (err error) {
	if !PathExist(path) {
		if err = MkFileLastDir(path); nil != nil {
			return err
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		if err = f.Close(); err != nil {
			return err
		}
	}
	return
}

func MkFileLastDir(path string) (err error) {
	if !PathExist(path) {
		f := func(c rune) bool {
			if c == '\\' || c == '/' {
				return true
			} else {
				return false
			}
		}

		index := strings.LastIndexFunc(path, f)
		if index != -1 {

			dirs := path[:index]
			err = os.MkdirAll(dirs, os.ModePerm)
			if err != nil {
				return
			}
		}
	}
	return
}
