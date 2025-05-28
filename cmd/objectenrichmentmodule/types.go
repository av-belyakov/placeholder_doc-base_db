package objectenrichmentmodule

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/storagetemporary"
)

// ObjectEnrichment занимается обогащением объектов дополнительной информацией
type ObjectEnrichment[T any] struct {
	storage           *storagetemporary.StorageTemporary[T]
	logger            interfaces.Logger
	ChInput, ChOutput chan ChanSetting
}

type ChanSetting struct {
	Data    []byte
	TaskId  string
	Command string
}
