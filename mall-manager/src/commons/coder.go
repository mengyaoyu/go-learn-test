package commons

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func EncoderByMd5(param string) string {
	hash := md5.New()
	hash.Write([]byte(param))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Base64Encoder(content string) string {
	return base64.URLEncoding.EncodeToString([]byte(content))
}
