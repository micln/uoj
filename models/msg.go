package models

import (
	"time"
)

/*
	- sb sloved a problem.
	- sb publish a article.
	-
*/
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
