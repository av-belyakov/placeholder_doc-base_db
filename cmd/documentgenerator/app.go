package documentgenerator

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

// New конструктор нового генератора документов
func New(counter *countermessage.CounterMessage, logger interfaces.Logger, chInput <-chan interfaces.CustomJsonDecoder) *DocumentGenerator {
	return &DocumentGenerator{
		counter: counter,
		logger:  logger,
		chInput: chInput,
	}
}
