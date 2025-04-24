package handlers

import (
	"strings"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewListHandlerAlertArtifacts обработчик событий 'alert.artifacts.*' типа для объекта 'alert'
func NewListHandlerAlertArtifacts(saa *SupportiveAlertArtifacts) map[string][]func(any) {
	return map[string][]func(any){
		//--- ioc ---
		"alert.artifacts.ioc": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.ioc",
				i,
				saa.GetArtifactTmp().SetAnyIoc,
			)
		}},
		//--- tlp ---
		"alert.artifacts.tlp": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.tlp",
				i,
				saa.GetArtifactTmp().SetAnyTlp,
			)
		}},
		//--- _id ---
		"alert.artifacts._id": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts._id",
				i,
				saa.GetArtifactTmp().SetAnyUnderliningId,
			)
		}},
		//--- id ---
		"alert.artifacts.id": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.id",
				i,
				saa.GetArtifactTmp().SetAnyId,
			)
		}},
		//--- _type ---
		"alert.artifacts._type": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts._type",
				i,
				saa.GetArtifactTmp().SetAnyUnderliningType,
			)
		}},
		//--- createdAt ---
		"alert.artifacts.createdAt": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.createdAt",
				i,
				saa.GetArtifactTmp().SetAnyCreatedAt,
			)
		}},
		//--- startDate ---
		"alert.artifacts.startDate": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.startDate",
				i,
				saa.GetArtifactTmp().SetAnyStartDate,
			)
		}},
		//--- createdBy ---
		"alert.artifacts.createdBy": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.createdBy",
				i,
				saa.GetArtifactTmp().SetAnyCreatedBy,
			)
		}},
		//--- data ---
		"alert.artifacts.data": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.data",
				i,
				saa.GetArtifactTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"alert.artifacts.dataType": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.dataType",
				i,
				saa.GetArtifactTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"alert.artifacts.message": {func(i any) {
			saa.HandlerValue(
				"alert.artifacts.message",
				i,
				saa.GetArtifactTmp().SetAnyMessage,
			)
		}},
		//--- tags ---
		"alert.artifacts.tags": {
			func(i any) {
				saa.HandlerValue(
					"alert.artifacts.tags",
					i,
					func(i any) {
						key, value := supportingfunctions.HandlerTag(i)
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
