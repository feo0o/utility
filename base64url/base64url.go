package base64url

import (
	"encoding/base64"
	"strings"
)

func Encode(src []byte) string {
	s := base64.StdEncoding.EncodeToString(src)
	s = strings.Replace(s, "+", "-", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "=", "", -1)
	return s
}

func Decode(str string) (d []byte, err error) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	switch len(str) % 4 {
	case 0:
	case 2:
		str += "=="
	case 3:
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
