package tokenbucket

import (
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestTokenBucketConcurrentAccess10User(t *testing.T) {
	rateLimitMultiIP(t, 10, 10, 2)
}

func TestTokenBucketConcurrentAccess100User(t *testing.T) {
	rateLimitMultiIP(t, 100, 1000, 2)
}

func TestTokenBucketConcurrentAccess1000User(t *testing.T) {
	rateLimitMultiIP(t, 1000, 1000, 2)
}

func rateLimitMultiIP(t *testing.T, goroutines int, ops int, period int) {
	wg := sync.WaitGroup{}

	tokenBucket := New(period)
	go tokenBucket.Run()

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 1; j <= ops; j++ {
				n := tokenBucket.Get("A" + strconv.Itoa(i))
				assert.Equal(t, j, n)
			}
		}(i)
	}
	wg.Wait()
}
