package model

import (
	"time"
)

type CurState struct {
	date     time.Time
	position Position
	event    Event
	ppls     []Person
}
