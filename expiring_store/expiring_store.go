package expiring_store

import (
	"crypto/sha1"
	"fmt"
	"io"
	r "math/rand"
	s "sync"
	t "time"
)

var (
	sha               = sha1.New()
	defaultExpiryTime = int64(5 * t.Minute)
)

type ExpiringStore struct {
	lock       s.RWMutex
	val        map[string]string
	createdAt  map[string]int64
	expiryTime int64
}

func (e *ExpiringStore) Set(value string) string {
	e.lock.Lock()
	defer e.lock.Unlock()
	key := generateKey()
	e.val[key] = value
	e.createdAt[key] = now() + e.expiryTime

	return key
}

func (e *ExpiringStore) Get(key string) (string, bool) {
	e.lock.RLock()
	defer e.lock.RUnlock()
	// FIXME needs err checking
	createdAt, _ := e.createdAt[key]
	if now() > createdAt {
		delete(e.createdAt, key)
		delete(e.val, key)
		return "", false
	}

	val, status := e.val[key]
	return val, status
}

func generateKey() string {
	io.WriteString(sha, fmt.Sprintf("%s", r.New(r.NewSource(now())).Int()))
	key := fmt.Sprintf("%x", sha.Sum(nil))

	return key
}

func now() int64 {
	return t.Now().UnixNano()
}

func New(expiryTime int64) *ExpiringStore {
	if expiryTime == 0 {
		expiryTime = defaultExpiryTime
	}
	return &ExpiringStore{
		val:        make(map[string]string),
		createdAt:  make(map[string]int64),
		expiryTime: expiryTime}
}
