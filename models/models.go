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

func Str2n(s string) (res int) {
	fmt.Sscanf(s, "%d", &res)
	return
}

func Sha120(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return string(h.Sum(nil))
}

func BadMail(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return true
	}
	return false
}

func HasNull(args ...string) bool {
	for _, v := range args {
		if len(v) == 0 {
			return true
		}
	}
	return false
}
