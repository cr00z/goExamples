package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	want := "3\n2\n1\nGo!"

	Countdown(buffer, spySleeper)
	got := buffer.String()

	assert.Equal(t, got, want)
	assert.Equal(t, spySleeper.Calls, 3)
}

func TestSleepBeforeEveryPrint(t *testing.T) {
	spySleepPrinter := &SpyCountdownOperations{}
	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	Countdown(spySleepPrinter, spySleepPrinter)

	assert.Equal(t, spySleepPrinter.Calls, want)
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	assert.Equal(t, spyTime.durationSlept, sleepTime)
}
