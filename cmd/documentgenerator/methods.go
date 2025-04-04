package documentgenerator

// NewVerifiedCase новый элемент содержащий проверенный объект типа 'case'
func NewVerifiedCase() *VerifiedCase {
	return &VerifiedCase{}
}

// GetID идентификатор объекта
func (vc *VerifiedCase) GetID() string {
	return vc.ID
}

// SetID идентификатор объекта
func (vc *VerifiedCase) SetID(v string) {
	vc.ID = v
}

// GetSource наименование источника
func (vc *VerifiedCase) GetSource() string {
	return vc.Source
}

// SetSource наименование источника
func (vc *VerifiedCase) SetSource(v string) {
	vc.Source = v
}
