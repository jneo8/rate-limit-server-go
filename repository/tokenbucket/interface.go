package tokenbucket

// TokenBucket ...
type TokenBucket interface {
	Visit(ip string) bool
	Supplement(t int64) error
	Run() error
}
