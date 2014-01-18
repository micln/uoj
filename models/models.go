package models

import (
	"crypto/sha1"
	"fmt"
	"io"
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
