package models

import (
	"time"
)

var (
	Msgs []*Msg
)

type Msg struct {
	Id    int
	Uid   string
	Ctx   string
	Time  time.Time
	Live  bool
	Tag   string
	Cnt   []int
	Quote int
	Type  int //
}
