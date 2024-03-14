package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	a := sig(time.Second * 3)
	b := sig(5 * time.Minute)
	c := sig(5 * time.Hour)

	d := Or(a, b, c)
	_, ok := <-d
	if ok {
		t.Fatalf("Канал не закрыт")
	}

}
