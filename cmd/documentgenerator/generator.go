package documentgenerator

// NewGenerateObjectNewFormat генератор новых объектов
func NewGenerateObjectsFormatMISP() *GenerateObjectsFormatMISP {
	return &GenerateObjectsFormatMISP{
		mispModule:    settings.MispModule,
		sqlite3Module: settings.Sqlite3Module,
		listRule:      settings.ListRule,
		counter:       settings.Counter,
		logger:        settings.Logger,
	}
}

func generatorNewFormat() {

}
