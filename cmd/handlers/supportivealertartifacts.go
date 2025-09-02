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
func (sa *SupportiveAlertArtifacts) GetArtifacts() map[string][]alertartifacts.Artifacts {
	sa.listAcceptedFields = []string(nil)

	if sa.currentKey != "" {
		_, _ = PostProcessingUserType(&sa.artifactTmp)
		sa.artifacts[sa.currentKey] = append(sa.artifacts[sa.currentKey], sa.artifactTmp)
	}

	sa.currentKey = ""
	sa.artifactTmp = *alertartifacts.NewArtifact()

	return sa.artifacts
}

// GetArtifactTmp возвращает временный объект artifact
func (sa *SupportiveAlertArtifacts) GetArtifactTmp() *alertartifacts.Artifacts {
	return &sa.artifactTmp
}

func (sa *SupportiveAlertArtifacts) HandlerValue(fieldBranch string, a any, f func(any)) {
	if fieldBranch == "alert.artifacts.dataType" {
		str := fmt.Sprint(a)
		if _, ok := sa.artifacts[str]; !ok {
			sa.artifacts[str] = []alertartifacts.Artifacts(nil)
		}

		if sa.isExistFieldBranch(fieldBranch) {
			sa.listAcceptedFields = []string(nil)

			_, _ = PostProcessingUserType(&sa.artifactTmp)
			sa.artifacts[sa.currentKey] = append(sa.artifacts[sa.currentKey], sa.artifactTmp)

			sa.artifactTmp = *alertartifacts.NewArtifact()
		}

		sa.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "alert.artifacts.tags" && sa.isExistFieldBranch(fieldBranch) {
		sa.listAcceptedFields = []string(nil)

		if _, ok := sa.artifacts[sa.currentKey]; !ok {
			sa.artifacts[sa.currentKey] = []alertartifacts.Artifacts(nil)
		}

		_, _ = PostProcessingUserType(&sa.artifactTmp)
		sa.artifacts[sa.currentKey] = append(sa.artifacts[sa.currentKey], sa.artifactTmp)

		sa.artifactTmp = *alertartifacts.NewArtifact()
	}

	sa.listAcceptedFields = append(sa.listAcceptedFields, fieldBranch)

	f(a)
}

func (sa *SupportiveAlertArtifacts) isExistFieldBranch(value string) bool {
	for _, v := range sa.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}
