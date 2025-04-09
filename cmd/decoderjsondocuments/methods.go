package decoderjsondocuments

// GetUUID уникальный идентификатор
func (settings *ChanInputSettings) GetUUID() string {
	return settings.UUID
}

// GetUUID уникальный идентификатор
func (settings *ChanInputSettings) SetUUID(v string) {
	settings.UUID = v
}

// GetFieldName имя поля
func (settings *ChanInputSettings) GetFieldName() string {
	return settings.FieldName
}

// SetFieldName имя поля
func (settings *ChanInputSettings) SetFieldName(v string) {
	settings.FieldName = v
}

// GetValue значение
func (settings *ChanInputSettings) GetValue() any {
	return settings.Value
}

// SetValue значение
func (settings *ChanInputSettings) SetValue(v any) {
	settings.Value = v
}

// GetValueType тип значения
func (settings *ChanInputSettings) GetValueType() string {
	return settings.ValueType
}

// SetValueType тип значения
func (settings *ChanInputSettings) SetValueType(v string) {
	settings.ValueType = v
}

// GetFieldBranch путь к полю
func (settings *ChanInputSettings) GetFieldBranch() string {
	return settings.FieldBranch
}

// SetFieldBranch путь к полю
func (settings *ChanInputSettings) SetFieldBranch(v string) {
	settings.FieldBranch = v
}
