package tokenbucket

// TokenBucket ...
type TokenBucket interface {
	Get(ip string) int
	Supplement(t int64) error
	Run() error
}
