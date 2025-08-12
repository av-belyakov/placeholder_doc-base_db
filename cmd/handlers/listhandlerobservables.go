package handlers

import "strings"

// NewListHandlerObservables обработчик событий типа 'observables.*' для объекта 'case'
func NewListHandlerObservables(so *SupportiveObservables) map[string][]func(any) {
	return map[string][]func(any){
		//--- ioc ---
		"observables.ioc": {func(a any) {
			so.HandlerValue(
				"observables.ioc",
				a,
				so.GetObservableTmp().SetAnyIoc,
			)
		}},
		//--- sighted ---
		"observables.sighted": {func(a any) {
			so.HandlerValue(
				"observables.sighted",
				a,
				so.GetObservableTmp().SetAnySighted,
			)
		}},
		//--- ignoreSimilarity ---
		"observables.ignoreSimilarity": {func(a any) {
			so.HandlerValue(
				"observables.ignoreSimilarity",
				a,
				so.GetObservableTmp().SetAnyIgnoreSimilarity,
			)
		}},
		//--- tlp ---
		"observables.tlp": {func(a any) {
			so.HandlerValue(
				"observables.tlp",
				a,
				so.GetObservableTmp().SetAnyTlp,
			)
		}},
		//--- _createdAt ---
		"observables._createdAt": {func(a any) {
			so.HandlerValue(
				"observables._createdAt",
				a,
				so.GetObservableTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _updatedAt ---
		"observables._updatedAt": {func(a any) {
			so.HandlerValue(
				"observables._updatedAt",
				a,
				so.GetObservableTmp().SetAnyUnderliningUpdatedAt,
			)
		}},
		//--- startDate ---
		"observables.startDate": {func(a any) {
			so.HandlerValue(
				"observables.startDate",
				a,
				so.GetObservableTmp().SetAnyStartDate,
			)
		}},
		//--- _createdBy ---
		"observables._createdBy": {func(a any) {
			so.HandlerValue(
				"observables._createdBy",
				a,
				so.GetObservableTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- _updatedBy ---
		"observables._updatedBy": {func(a any) {
			so.HandlerValue(
				"observables._updatedBy",
				a,
				so.GetObservableTmp().SetAnyUnderliningUpdatedBy,
			)
		}},
		//--- _id ---
		"observables._id": {func(a any) {
			so.HandlerValue(
				"observables._id",
				a,
				so.GetObservableTmp().SetAnyUnderliningId,
			)
		}},
		//--- _type ---
		"observables._type": {func(a any) {
			so.HandlerValue(
				"observables._type",
				a,
				so.GetObservableTmp().SetAnyUnderliningType,
			)
		}},
		//--- data ---
		"observables.data": {func(a any) {
			so.HandlerValue(
				"observables.data",
				a,
				so.GetObservableTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"observables.dataType": {func(a any) {
			so.HandlerValue(
				"observables.dataType",
				a,
				so.GetObservableTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"observables.message": {func(a any) {
			so.HandlerValue(
				"observables.message",
				a,
				so.GetObservableTmp().SetAnyMessage,
			)
		}},

		//--- tags ---
		"observables.tags": {
			func(a any) {
				so.HandlerValue(
					"observables.tags",
					a,
					func(a any) {
						key, value := HandlerTag(a)
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
		"observables.attachment.id": {func(a any) {
			so.HandlerValue(
				"observables.attachment.id",
				a,
				so.GetObservableTmp().Attachment.SetAnyId,
			)
		}},
		//--- attachment.size ---
		"observables.attachment.size": {func(a any) {
			so.HandlerValue(
				"observables.attachment.size",
				a,
				so.GetObservableTmp().Attachment.SetAnySize,
			)
		}},
		// --- attachment.name ---
		"observables.attachment.name": {func(a any) {
			so.HandlerValue(
				"observables.attachment.name",
				a,
				so.GetObservableTmp().Attachment.SetAnyName,
			)
		}},
		// --- attachment.contentType ---
		"observables.attachment.contentType": {func(a any) {
			so.HandlerValue(
				"observables.attachment.contentType",
				a,
				so.GetObservableTmp().Attachment.SetAnyContentType,
			)
		}},
		// --- attachment.hashes ---
		"observables.attachment.hashes": {func(a any) {
			so.HandlerValue(
				"observables.attachment.hashes",
				a,
				so.GetObservableTmp().Attachment.SetAnyHashes,
			)
		}},
	}
}
