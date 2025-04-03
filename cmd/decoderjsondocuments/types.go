package decoderjsondocuments

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

type DecoderJsonDocuments struct {
	logger  interfaces.Logger
	counter *countermessage.CounterMessage
}

// ChanInputSettings параметры канала декодера JSON документов
type ChanInputSettings struct {
	Value       any    //любые передаваемые данные
	UUID        string //уникальный идентификатор в формате UUID
	FieldName   string //наименование обработанного поля
	ValueType   string //тип переданного значения (string, int и т.д.)
	FieldBranch string //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
}
