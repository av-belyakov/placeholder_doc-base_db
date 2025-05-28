package storagetemporary

import (
	"context"
	"errors"
	"time"
)

// New инициирует новое хранилище
func New[T any](ctx context.Context, opts ...OptionsStorageTemporary[T]) (st *StorageTemporary[T], err error) {
	st = &StorageTemporary[T]{
		cache:      map[string]storage[T]{},
		timeTick:   5 * time.Second,
		timeToLive: 30 * time.Second,
	}

	for _, opt := range opts {
		if err = opt(st); err != nil {
			return
		}
	}

	go st.deleteOldIncomingRequests(ctx)

	return
}

// WithTimeTick интервал проверки данных в кеше, в секундах от 2 до 15
func WithTimeTick[T any](v int) OptionsStorageTemporary[T] {
	return func(st *StorageTemporary[T]) error {
		if v < 2 || v > 15 {
			return errors.New("the time tick interval of a cache entry should be between 2 and 15 seconds")
		}

		st.timeTick = time.Duration(v) * time.Second

		return nil
	}
}

// WithTimeToLive время жизни для объекта в кеше, в секундах от 3 до 3600
func WithTimeToLive[T any](v int) OptionsStorageTemporary[T] {
	return func(st *StorageTemporary[T]) error {
		if v < 3 || v > 3600 {
			return errors.New("the lifetime of a cache entry should be between 3 and 3600 seconds")
		}

		st.timeToLive = time.Duration(v) * time.Second

		return nil
	}
}
