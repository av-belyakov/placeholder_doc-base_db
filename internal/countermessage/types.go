package countermessage

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/informationcountingstorage"
)

// CounterMessage счетчик сообщений
type CounterMessage struct {
	storage  *informationcountingstorage.InformationCountingStorage
	chInput  chan DataCounterMessage
	chOutput chan<- interfaces.Messager
}

// SomeMessage некоторое сообщение
type SomeMessage struct {
	Type, Message string
}

// DataCounterMessage содержит информацию для подсчета
type DataCounterMessage struct {
	msgType string
	count   int
}
