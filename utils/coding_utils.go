package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/satori/go.uuid"
)

func GetMD5Coding(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GetUUId() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
