package tokenbucket

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type repo struct {
	Bucket           *sync.Map
	Logger           *log.Logger
	SupplementBucket map[int64][]string
	RWMutex          *sync.RWMutex
}

func (r *repo) Run() error {
	ticker := time.NewTicker(1 * time.Second).C
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for t := range ticker {
			r.Supplement(t.Add(-10 * time.Second).Unix())
		}
	}()
	wg.Wait()
	return nil
}

// Visit return true if ip rate limit < settings limit.
func (r *repo) Visit(ip string) bool {
	r.increment(ip)
	v, ok := r.Bucket.LoadOrStore(ip, 1)
	// First time
	if !ok {
		return true
	}

	n, ok := v.(int)
	// If type error, reset rate limit for ip.
	if !ok {
		r.Bucket.Delete(ip)
	}

	if n >= 60 {
		return false
	}
	r.Bucket.Store(ip, n+1)
	return true
}

// Supplement token to bucket.
func (r *repo) Supplement(t int64) error {
	ips, ok := r.SupplementBucket[t]
	if !ok {
		return nil
	}

	for _, ip := range ips {
		r.decreaseBucket(ip)
	}
	return nil
}

func (r *repo) decreaseBucket(ip string) {
	if v, ok := r.Bucket.Load(ip); ok {
		if v == 1 {
			r.Bucket.Delete(v)
		}
		if n, typeOk := v.(int); typeOk {
			r.Bucket.Store(ip, n-1)
		}
	}
}

func (r *repo) increment(ip string) {
	r.RWMutex.Lock()
	defer r.RWMutex.Unlock()
	now := time.Now().Unix()
	_, ok := r.SupplementBucket[now]
	if !ok {
		r.SupplementBucket[now] = []string{ip}
	}
	r.SupplementBucket[now] = append(
		r.SupplementBucket[now],
		ip,
	)
}
