package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	want := "3\n2\n1\nGo!"
	buffer := &bytes.Buffer{}
	spy := &SpySleeper{}

	Countdown(buffer, spy)
	got := buffer.String()

	assert.Equal(t, got, want)
	assert.Equal(t, spy.Calls, 3)
}

func TestSleepBeforeEveryPrint(t *testing.T) {
	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}
	spySleepPrinter := &SpyCountdownOperations{}

	Countdown(spySleepPrinter, spySleepPrinter)

	assert.Equal(t, spySleepPrinter.Calls, want)
	fmt.Println(spySleepPrinter.Calls)
}
