package entities

import "time"

type Exchange struct {
	feeds map[int][]LiveFeed
}

type LiveFeed struct {
	InstrumentId    int
	LastTradedPrice float64
	TimeStamp       time.Time
}
