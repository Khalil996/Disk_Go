package common

import (
	"crypto/md5"
	"fmt"
)

func MD5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
