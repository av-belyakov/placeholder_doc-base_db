package decoderjsondocuments

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

type DecoderJsonDocuments struct {
	logger  interfaces.Logger
	counter interfaces.Counter
}

// ChanInputSettings параметры канала декодера JSON документов
type ChanInputSettings struct {
	Value       any    //любые передаваемые данные
	UUID        string //уникальный идентификатор в формате UUID
	FieldName   string //наименование обработанного поля
	ValueType   string //тип переданного значения (string, int и т.д.)
	FieldBranch string //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
}
