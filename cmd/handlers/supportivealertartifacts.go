package handlers

import (
	"fmt"

	alertartifacts "github.com/av-belyakov/objectsthehiveformat/alertartifacts"
)

// SupportiveAlertArtifacts вспомогательный тип для для обработки alert.artifacts
type SupportiveAlertArtifacts struct {
	artifacts          map[string][]alertartifacts.Artifacts
	artifactTmp        alertartifacts.Artifacts
	listAcceptedFields []string
	currentKey         string
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа alert.artifacts
func NewSupportiveAlertArtifacts() *SupportiveAlertArtifacts {
	return &SupportiveAlertArtifacts{
		listAcceptedFields: []string(nil),
		artifactTmp:        *alertartifacts.NewArtifact(),
		artifacts:          make(map[string][]alertartifacts.Artifacts),
	}
}

// GetArtifacts возвращает map[string][]datamodels.AlertArtifact, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из a.artifactTmp в
// a.artifacts, так как artifacts автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются artifactTmp
func (a *SupportiveAlertArtifacts) GetArtifacts() map[string][]alertartifacts.Artifacts {
	a.listAcceptedFields = []string(nil)

	if a.currentKey != "" {
		_, _ = PostProcessingUserType[*alertartifacts.Artifacts](&a.artifactTmp)
		a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)
	}

	a.currentKey = ""
	a.artifactTmp = *alertartifacts.NewArtifact()

	return a.artifacts
}

// GetArtifactTmp возвращает временный объект artifact
func (a *SupportiveAlertArtifacts) GetArtifactTmp() *alertartifacts.Artifacts {
	return &a.artifactTmp
}

func (a *SupportiveAlertArtifacts) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	if fieldBranch == "alert.artifacts.dataType" {
		str := fmt.Sprint(i)
		if _, ok := a.artifacts[str]; !ok {
			a.artifacts[str] = []alertartifacts.Artifacts(nil)
		}

		if a.isExistFieldBranch(fieldBranch) {
			a.listAcceptedFields = []string(nil)

			_, _ = PostProcessingUserType[*alertartifacts.Artifacts](&a.artifactTmp)
			a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)

			a.artifactTmp = *alertartifacts.NewArtifact()
		}

		a.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "alert.artifacts.tags" && a.isExistFieldBranch(fieldBranch) {
		a.listAcceptedFields = []string(nil)

		if _, ok := a.artifacts[a.currentKey]; !ok {
			a.artifacts[a.currentKey] = []alertartifacts.Artifacts(nil)
		}

		_, _ = PostProcessingUserType[*alertartifacts.Artifacts](&a.artifactTmp)
		a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)

		a.artifactTmp = *alertartifacts.NewArtifact()
	}

	a.listAcceptedFields = append(a.listAcceptedFields, fieldBranch)

	f(i)
}

func (a *SupportiveAlertArtifacts) isExistFieldBranch(value string) bool {
	for _, v := range a.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

/*
// ArtifactForAlert содержит артефакт для типа 'alert'
type ArtifactForAlert struct {
	Tags           map[string][]string `json:"tags" bson:"tags"`                               //теги после обработки
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`             //список snort сигнатур (строка)
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`                         //все теги
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"` //список snort сигнатур (число)
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`             //сенсор id
	common.CommonArtifactType
}
*/
