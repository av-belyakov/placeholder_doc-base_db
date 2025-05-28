package objectenrichmentmodule

import (
	"context"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/storagetemporary"
)

func New[T any](ctx context.Context, logging interfaces.Logger) *ObjectEnrichment[T] {
	s, _ := storagetemporary.New(
		ctx,
		storagetemporary.WithTimeTick[T](3),
		storagetemporary.WithTimeToLive[T](60),
	)

	e := &ObjectEnrichment[T]{
		logger:   logging,
		storage:  s,
		ChInput:  make(chan ChanSetting),
		ChOutput: make(chan ChanSetting),
	}

	go e.handlerRequest(ctx)

	return e
}

// handlerRequest обрабатывает входящие запросы
func (e *ObjectEnrichment[T]) handlerRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case data := <-e.ChInput:
			//тут обрабатываем входящие запросы на обогащение данными

		}
	}
}
