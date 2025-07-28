package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		leavingOrder = []string{}
		dine()

		if len(leavingOrder) != 5 {
			t.Errorf("incorrect length of slice")
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"Half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		leavingOrder = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()

		if len(leavingOrder) != 5 {
			t.Errorf("%s: incorrect length of slice", e.name)
		}
	}
}
