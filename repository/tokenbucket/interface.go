package tokenbucket

// TokenBucket is a an algorithm used in packet switched computer networks and telecommunications networks.
// Only define interface here.
type TokenBucket interface {
	// Get return ip's connection number
	Get(ip string) int
	// Supplement token by timestamp
	Supplement(t int64) error
	// Run, start the token bucket
	Run() error
}
