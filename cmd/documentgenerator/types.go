package documentgenerator

// ChanInputCreateNewFormat входные данные для канала приёма информации
type ChanInputCreateNewFormat struct {
	Value       any    //любые передаваемые данные
	UUID        string //уникальный идентификатор в формате UUID
	FieldName   string //наименование поля
	ValueType   string //тип передаваемого значения (string, int и т.д.)
	FieldBranch string //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
}
