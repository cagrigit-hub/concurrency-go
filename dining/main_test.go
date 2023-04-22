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
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("Expected 5, got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Millisecond * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, tt := range theTests {
		orderFinished = []string{}
		eatTime = tt.delay
		thinkTime = tt.delay
		sleepTime = tt.delay
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("Expected 5, got %d", len(orderFinished))
		}
	}
}
