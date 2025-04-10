package handlers

import (
	common "github.com/av-belyakov/objectsthehiveformat/common"
)

// SupportiveAlertArtifacts вспомогательный тип для для обработки alert.artifacts
type SupportiveAlertArtifacts struct {
	artifacts          map[string][]ArtifactForAlert
	artifactTmp        ArtifactForAlert
	listAcceptedFields []string
	currentKey         string
}

// ArtifactForAlert содержит артефакт для типа 'alert'
type ArtifactForAlert struct {
	Tags           map[string][]string `json:"tags" bson:"tags"`                               //теги после обработки
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`             //список snort сигнатур (строка)
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`                         //все теги
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"` //список snort сигнатур (число)
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`             //сенсор id
	common.CommonArtifactType
}
