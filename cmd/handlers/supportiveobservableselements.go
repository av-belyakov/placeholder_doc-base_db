package handlers

import (
	"fmt"

	"slices"

	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// SupportiveObservables вспомогательный тип для для обработки observables
type SupportiveObservables struct {
	currentKey                string
	listAcceptedFields        []string
	listAcceptedFieldsReports []string
	observableTmp             caseobservables.Observable
	observables               map[string][]caseobservables.Observable
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа observables
func NewSupportiveObservables() *SupportiveObservables {
	return &SupportiveObservables{
		listAcceptedFields:        []string(nil),
		listAcceptedFieldsReports: []string(nil),
		observableTmp:             *caseobservables.NewObservable(),
		observables:               make(map[string][]caseobservables.Observable),
	}
}

// GetObservables возвращает map[string][]datamodels.ObservableMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из o.observableTmp в
// o.observables, так как observables автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются observableTmp
func (o *SupportiveObservables) GetObservables() map[string][]caseobservables.Observable {
	o.listAcceptedFields = []string(nil)

	if o.currentKey != "" {
		_, _ = supportingfunctions.PostProcessingUserType(&o.observableTmp)
		o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)
	}

	o.currentKey = ""
	o.observableTmp = *caseobservables.NewObservable()

	return o.observables
}

// GetObservableTmp возвращает временный объект observable
func (o *SupportiveObservables) GetObservableTmp() *caseobservables.Observable {
	return &o.observableTmp
}

func (o *SupportiveObservables) HandlerValue(fieldBranch string, a any, f func(any)) {
	if fieldBranch == "observables.dataType" {
		str := fmt.Sprint(a)
		if _, ok := o.observables[str]; !ok {
			o.observables[str] = []caseobservables.Observable(nil)
		}

		if o.isExistFieldBranch(fieldBranch) {
			o.listAcceptedFields = []string(nil)
			_, _ = supportingfunctions.PostProcessingUserType(&o.observableTmp)
			o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)

			o.observableTmp = *caseobservables.NewObservable()
		}

		o.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "observables.tags" && fieldBranch != "observables.attachment.hashes" && o.isExistFieldBranch((fieldBranch)) {
		o.listAcceptedFields = []string(nil)

		if _, ok := o.observables[o.currentKey]; !ok {
			o.observables[o.currentKey] = []caseobservables.Observable(nil)
		}

		_, _ = supportingfunctions.PostProcessingUserType(&o.observableTmp)
		o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)

		o.observableTmp = *caseobservables.NewObservable()
	}

	o.listAcceptedFields = append(o.listAcceptedFields, fieldBranch)

	f(a)
}

func (o *SupportiveObservables) isExistFieldBranch(value string) bool {
	return slices.Contains(o.listAcceptedFields, value)
}
