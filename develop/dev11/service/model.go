package service

import (
	"sync"
	"time"
)

type Event struct {
	Date time.Time
	Mes  string
}

type Calendar struct {
	m     *sync.Mutex
	store map[int]Event
}
