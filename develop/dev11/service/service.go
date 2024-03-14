package service

import (
	"errors"
	"reflect"
	"strconv"
	"sync"
	_ "sync"
	"time"
)

func NewStore() *Calendar {
	return &Calendar{
		m:     &sync.Mutex{},
		store: make(map[int]Event),
	}
}

func (c *Calendar) CreateEvent(date time.Time, mes string) int {
	event := Event{date, mes}

	c.m.Lock()
	defer c.m.Unlock()

	id := len(c.store)

	for {
		if reflect.DeepEqual(c.store[id], Event{}) {
			c.store[id] = event
			return id
		}
		id++
	}
}

func (c *Calendar) UpdateEvent(id int, date time.Time, mes string) (Event, error) {
	c.m.Lock()
	defer c.m.Unlock()

	//возвращаем ошибку если элемента нет
	if reflect.DeepEqual(c.store[id], Event{}) {
		return Event{}, errors.New("503: invalid element")
	}

	event := Event{date, mes}

	c.store[id] = event

	return c.store[id], nil
}

func (c *Calendar) DeleteEvent(id int) error {
	c.m.Lock()
	defer c.m.Unlock()

	if reflect.DeepEqual(c.store[id], Event{}) {
		return errors.New("503: No event for delete")
	}

	delete(c.store, id)

	return nil
}

func (c *Calendar) EventsForDate(date time.Time, days int) ([]Event, error) {
	var result []Event
	duration, err := time.ParseDuration(strconv.Itoa(days*24) + "h")
	if err != nil {
		return []Event{}, errors.New("503: Invalid duration")
	}
	for _, event := range c.store {
		if event.Date.After(date.AddDate(0, 0, -1)) && event.Date.Before(date.Add(duration)) {
			result = append(result, event)
		}
	}

	return result, nil
}
