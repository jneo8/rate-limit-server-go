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
		n := tokenBucket.Get("A")
		if i < 60 && n > 60 {
			t.Error("Rate limit not work when access < 60")
		}
		if i >= 60 && n < 60 {
			t.Error("Rate limit not work when access > 60")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		if n := tokenBucket.Get("A"); n < 60 {
			t.Error("Rate limit not work after 2 secs")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 11)
		if n := tokenBucket.Get("A"); n > 60 {
			t.Error("Rate limit not work after 10 secs")
		}
	}()
	wg.Wait()
}
