package models

import (
	"time"
)

type Task struct {
	Id     int
	Author string
	Ctx    string
	Time   time.Time
	Live   bool
	Tag    string
	Cnt    []int
	Quote  int
}
