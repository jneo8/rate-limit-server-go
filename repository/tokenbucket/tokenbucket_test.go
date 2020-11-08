package tokenbucket

import (
	"sync"
	"testing"
	"time"
)

// TestRateLimit work.
func TestRateLimitWork(t *testing.T) {
	tokenBucket := New()
	go tokenBucket.Run()
	for i := 0; i <= 120; i++ {
		pass := tokenBucket.Visit("A")
		if i < 60 && !pass {
			t.Error("Rate limit not work when access < 60")
		}
		if i >= 60 && pass {
			t.Error("Rate limit not work when access > 60")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		if pass := tokenBucket.Visit("A"); pass {
			t.Error("Rate limit not work after 2 secs")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 11)
		if pass := tokenBucket.Visit("A"); !pass {
			t.Error("Rate limit not work after 10 secs")
		}
	}()
	wg.Wait()
}
