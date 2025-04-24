package handlers

import "strings"

// NewListHandlerObservables обработчик событий типа 'observables.*' для объекта 'case'
func NewListHandlerObservables(so *SupportiveObservables) map[string][]func(any) {
	return map[string][]func(any){
		//--- ioc ---
		"observables.ioc": {func(i any) {
			so.HandlerValue(
				"observables.ioc",
				i,
				so.GetObservableTmp().SetAnyIoc,
			)
		}},
		//--- sighted ---
		"observables.sighted": {func(i any) {
			so.HandlerValue(
				"observables.sighted",
				i,
				so.GetObservableTmp().SetAnySighted,
			)
		}},
		//--- ignoreSimilarity ---
		"observables.ignoreSimilarity": {func(i any) {
			so.HandlerValue(
				"observables.ignoreSimilarity",
				i,
				so.GetObservableTmp().SetAnyIgnoreSimilarity,
			)
		}},
		//--- tlp ---
		"observables.tlp": {func(i any) {
			so.HandlerValue(
				"observables.tlp",
				i,
				so.GetObservableTmp().SetAnyTlp,
			)
		}},
		//--- _createdAt ---
		"observables._createdAt": {func(i any) {
			so.HandlerValue(
				"observables._createdAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _updatedAt ---
		"observables._updatedAt": {func(i any) {
			so.HandlerValue(
				"observables._updatedAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedAt,
			)
		}},
		//--- startDate ---
		"observables.startDate": {func(i any) {
			so.HandlerValue(
				"observables.startDate",
				i,
				so.GetObservableTmp().SetAnyStartDate,
			)
		}},
		//--- _createdBy ---
		"observables._createdBy": {func(i any) {
			so.HandlerValue(
				"observables._createdBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- _updatedBy ---
		"observables._updatedBy": {func(i any) {
			so.HandlerValue(
				"observables._updatedBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedBy,
			)
		}},
		//--- _id ---
		"observables._id": {func(i any) {
			so.HandlerValue(
				"observables._id",
				i,
				so.GetObservableTmp().SetAnyUnderliningId,
			)
		}},
		//--- _type ---
		"observables._type": {func(i any) {
			so.HandlerValue(
				"observables._type",
				i,
				so.GetObservableTmp().SetAnyUnderliningType,
			)
		}},
		//--- data ---
		"observables.data": {func(i any) {
			so.HandlerValue(
				"observables.data",
				i,
				so.GetObservableTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"observables.dataType": {func(i any) {
			so.HandlerValue(
				"observables.dataType",
				i,
				so.GetObservableTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"observables.message": {func(i any) {
			so.HandlerValue(
				"observables.message",
				i,
				so.GetObservableTmp().SetAnyMessage,
			)
		}},

		//--- tags ---
		"observables.tags": {
			func(i any) {
				so.HandlerValue(
					"observables.tags",
					i,
					func(i any) {
						key, value := HandlerTag(i)
						if value == "" {
							return
						}

						value = strings.TrimSpace(value)
						value = strings.Trim(value, "\"")
						so.GetObservableTmp().SetAnyTags(key, value)
					},
				)
			},
			so.GetObservableTmp().SetAnyTagsAll,
		},
		//--- attachment.id ---
		"observables.attachment.id": {func(i any) {
			so.HandlerValue(
				"observables.attachment.id",
				i,
				so.GetObservableTmp().Attachment.SetAnyId,
			)
		}},
		//--- attachment.size ---
		"observables.attachment.size": {func(i any) {
			so.HandlerValue(
				"observables.attachment.size",
				i,
				so.GetObservableTmp().Attachment.SetAnySize,
			)
		}},
		// --- attachment.name ---
		"observables.attachment.name": {func(i any) {
			so.HandlerValue(
				"observables.attachment.name",
				i,
				so.GetObservableTmp().Attachment.SetAnyName,
			)
		}},
		// --- attachment.contentType ---
		"observables.attachment.contentType": {func(i any) {
			so.HandlerValue(
				"observables.attachment.contentType",
				i,
				so.GetObservableTmp().Attachment.SetAnyContentType,
			)
		}},
		// --- attachment.hashes ---
		"observables.attachment.hashes": {func(i any) {
			so.HandlerValue(
				"observables.attachment.hashes",
				i,
				so.GetObservableTmp().Attachment.SetAnyHashes,
			)
		}},
	}
}
