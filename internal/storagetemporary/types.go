package storagetemporary

import (
	"sync"
	"time"
)

// StorageTemporary временное хранилище
type StorageTemporary[T any] struct {
	mux        sync.RWMutex
	cache      map[string]storage[T]
	timeTick   time.Duration
	timeToLive time.Duration
}

type storage[T any] struct {
	timeExpiry time.Time
	data       T
}

type OptionsStorageTemporary[T any] func(*StorageTemporary[T]) error
