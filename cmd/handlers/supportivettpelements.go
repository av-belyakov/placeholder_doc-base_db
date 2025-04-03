package handlers

import (
	"slices"

	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
)

var fieldsRepresentedAsList []string = []string{
	"ttp.extraData.pattern.tactics",
	"ttp.extraData.pattern.platforms",
	"ttp.extraData.pattern.dataSources",
	"ttp.extraData.pattern.systemRequirements",
	"ttp.extraData.pattern.permissionsRequired",
	"ttp.extraData.patternParent.tactics",
	"ttp.extraData.patternParent.platforms",
	"ttp.extraData.patternParent.dataSources",
	"ttp.extraData.patternParent.systemRequirements",
	"ttp.extraData.patternParent.permissionsRequired",
}

func isExistFieldsRepresentedAsList(field string, list []string) bool {
	return slices.Contains(list, field)
}

// SupportiveTtp вспомогательный тип для для обработки ttp
type SupportiveTtp struct {
	listAcceptedFields []string
	ttpTmp             casettps.Ttp
	ttps               []casettps.Ttp
}

// NewSupportiveTtp формирует вспомогательный объект для обработки thehive объектов типа ttp
func NewSupportiveTtp() *SupportiveTtp {
	return &SupportiveTtp{
		listAcceptedFields: []string(nil),
		ttpTmp:             *casettps.NewTtp(),
		ttps:               []casettps.Ttp(nil),
	}
}

// GetTtps возвращает []datamodels.TtpMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из sttp.ttpTmp в
// список sttp.ttps, так как ttps автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются sttp.ttpTmp
func (sttp *SupportiveTtp) GetTtps() []casettps.Ttp {
	sttp.listAcceptedFields = []string(nil)
	if sttp.ttpTmp.PatternId != "" {
		sttp.ttps = append(sttp.ttps, sttp.ttpTmp)
	}

	return sttp.ttps
}

// GetTtpTmp возвращает временный объект ttpTmp
func (sttp *SupportiveTtp) GetTtpTmp() *casettps.Ttp {
	return &sttp.ttpTmp
}

func (sttp *SupportiveTtp) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsRepresentedAsList)

	if !isExist && sttp.isExistFieldBranch(fieldBranch) {
		sttp.listAcceptedFields = []string(nil)
		sttp.ttps = append(sttp.ttps, sttp.ttpTmp)
		sttp.ttpTmp = *casettps.NewTtp()
	}

	sttp.listAcceptedFields = append(sttp.listAcceptedFields, fieldBranch)

	f(i)
}

func (sttp SupportiveTtp) isExistFieldBranch(value string) bool {
	for _, v := range sttp.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}
