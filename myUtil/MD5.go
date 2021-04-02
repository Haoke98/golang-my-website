package myUtil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(str string) string {
	m5 := md5.New()
	_, err := m5.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	md5String := hex.EncodeToString(m5.Sum(nil))
	return md5String
}
