package syncer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	counter.Inc()

	assert.Equal(t, 3, counter.Value())
}

func TestCounter1000Times(t *testing.T) {
	limit := 1000
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, limit, counter.Value())
}
