package encryptutil

import (
	"bufio"
	md52 "crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func MD5(file *multipart.FileHeader) string {
	f, err := file.Open()
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(f)
	md5 := md52.New()
	io.Copy(md5, reader)
	return hex.EncodeToString(md5.Sum(nil))
}
