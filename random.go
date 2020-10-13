package utility

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// RandomStr return a random string with given length that only contains lower/upper letters
func RandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// RandomInt return a pseudorandom int between min and max
func RandomInt(min, max int) (randInt int, err error) {
	if min > max {
		return 0, errors.New("wrong min/max sequence")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%(max-min+1) + min, nil
}
