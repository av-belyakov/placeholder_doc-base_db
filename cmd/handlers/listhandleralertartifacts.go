package handlers

import (
	"strings"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewListHandlerAlertArtifacts обработчик событий 'alert.artifacts.*' типа для объекта 'alert'
func NewListHandlerAlertArtifacts(saa *SupportiveAlertArtifacts) map[string][]func(any) {
	return map[string][]func(any){
		//--- ioc ---
		"alert.artifacts.ioc": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.ioc",
				a,
				saa.GetArtifactTmp().SetAnyIoc,
			)
		}},
		//--- tlp ---
		"alert.artifacts.tlp": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.tlp",
				a,
				saa.GetArtifactTmp().SetAnyTlp,
			)
		}},
		//--- _id ---
		"alert.artifacts._id": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts._id",
				a,
				saa.GetArtifactTmp().SetAnyUnderliningId,
			)
		}},
		//--- id ---
		"alert.artifacts.id": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.id",
				a,
				saa.GetArtifactTmp().SetAnyId,
			)
		}},
		//--- _type ---
		"alert.artifacts._type": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts._type",
				a,
				saa.GetArtifactTmp().SetAnyUnderliningType,
			)
		}},
		//--- createdAt ---
		"alert.artifacts.createdAt": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.createdAt",
				a,
				saa.GetArtifactTmp().SetAnyCreatedAt,
			)
		}},
		//--- startDate ---
		"alert.artifacts.startDate": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.startDate",
				a,
				saa.GetArtifactTmp().SetAnyStartDate,
			)
		}},
		//--- createdBy ---
		"alert.artifacts.createdBy": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.createdBy",
				a,
				saa.GetArtifactTmp().SetAnyCreatedBy,
			)
		}},
		//--- data ---
		"alert.artifacts.data": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.data",
				a,
				saa.GetArtifactTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"alert.artifacts.dataType": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.dataType",
				a,
				saa.GetArtifactTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"alert.artifacts.message": {func(a any) {
			saa.HandlerValue(
				"alert.artifacts.message",
				a,
				saa.GetArtifactTmp().SetAnyMessage,
			)
		}},
		//--- tags ---
		"alert.artifacts.tags": {
			func(a any) {
				saa.HandlerValue(
					"alert.artifacts.tags",
					a,
					func(a any) {
						key, value := supportingfunctions.HandlerTag(a)
						if value == "" {
							return
						}

						value = strings.TrimSpace(value)
						value = strings.Trim(value, "\"")
						saa.GetArtifactTmp().SetAnyTags(key, value)
					},
				)
			},
			saa.GetArtifactTmp().SetAnyTagsAll},
	}
}
