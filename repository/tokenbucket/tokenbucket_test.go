package tokenbucket

import (
	"sync"
	"testing"
	"time"
)

// TestRateLimit work.
func TestRateLimitWork(t *testing.T) {
	tokenBucket := New(5)

	go tokenBucket.Run()
	for i := 0; i <= 200; i++ {
		n := tokenBucket.Get("A")
		if i < 60 && n > 60 {
			t.Error("Rate limit not work when access not hit the limit")
		}
		if i >= 60 && n < 60 {
			t.Error("Rate limit not work when access hit the limit")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		if n := tokenBucket.Get("A"); n < 60 {
			t.Error("Rate limit not work before time period")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 6)
		if n := tokenBucket.Get("A"); n > 60 {
			t.Error("Rate limit not work after time period")
		}
	}()
	wg.Wait()
}
