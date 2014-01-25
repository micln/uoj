package models

import (
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
)

// Option Response
type Ops struct {
	OK   bool
	Info interface{}
}

// get a number in string
func Str2n(s string) (res int) {
	fmt.Sscanf(s, "%d", &res)
	return
}

// encode string to 20 bytes ciphertext by sha1
func Sha120(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return string(h.Sum(nil))
}

// check mail
func BadMail(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return true
	}
	return false
}

// return true if someone is nil
func HasNull(args ...string) bool {
	for _, v := range args {
		if len(v) == 0 {
			return true
		}
	}
	return false
}
