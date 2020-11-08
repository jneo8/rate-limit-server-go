package tokenbucket

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

// New return new TokenBucket.
func New() TokenBucket {
	logger := log.New()
	return &repo{
		Bucket:           &sync.Map{},
		SupplementBucket: make(map[int64][]string),
		RWMutex:          &sync.RWMutex{},
		Logger:           logger,
	}
}
