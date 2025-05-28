package storagetemporary

import (
	"context"
	"time"
)

// GetData данные в кеше
func (st *StorageTemporary[T]) GetData(key string) (T, bool) {
	st.mux.RLock()
	defer st.mux.RUnlock()

	if storage, ok := st.cache[key]; ok {
		return storage.data, ok
	}

	var zero T
	return zero, false
}

// SetData данные в кеше
func (st *StorageTemporary[T]) SetData(key string, value T) {
	st.mux.Lock()
	defer st.mux.Unlock()

	st.cache[key] = storage[T]{
		timeExpiry: time.Now().Add(st.timeToLive),
		data:       value,
	}
}

// DelData удалить данные из кеша
func (st *StorageTemporary[T]) DelData(key string) {
	st.mux.Lock()
	defer st.mux.Unlock()

	delete(st.cache, key)
}

// Cancel выполняет очистку кеша
func (st *StorageTemporary[T]) Cancel() {
	st.mux.Lock()
	defer st.mux.Unlock()

	st.cache = map[string]storage[T]{}
}

// DataSize размер кеша
func (st *StorageTemporary[T]) DataSize() int {
	return len(st.cache)
}

func (st *StorageTemporary[T]) deleteOldIncomingRequests(ctx context.Context) {
	go func(ctx context.Context) {
		tick := time.NewTicker(st.timeTick)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-tick.C:
				st.mux.Lock()

				for key, value := range st.cache {
					if value.timeExpiry.Before(time.Now()) {
						delete(st.cache, key)
					}
				}

				st.mux.Unlock()
			}
		}
	}(ctx)
}
