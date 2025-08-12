package handlers

// NewListHandlerTtp обработчик событий типа 'ttps.*' для объекта 'case'
func NewListHandlerTtp(sttp *SupportiveTtp) map[string][]func(any) {
	return map[string][]func(any){
		//--- occurDate ---
		"ttp.occurDate": {func(a any) {
			sttp.HandlerValue(
				"ttp.occurDate",
				a,
				sttp.GetTtpTmp().SetAnyOccurDate,
			)
		}},
		//--- _createdAt ---
		"ttp._createdAt": {func(a any) {
			sttp.HandlerValue(
				"ttp._createdAt",
				a,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _id ---
		"ttp._id": {func(a any) {
			sttp.HandlerValue(
				"ttp._id",
				a,
				sttp.GetTtpTmp().SetAnyUnderliningId,
			)
		}},
		//--- _createdBy ---
		"ttp._createdBy": {func(a any) {
			sttp.HandlerValue(
				"ttp._createdBy",
				a,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- patternId ---
		"ttp.patternId": {func(a any) {
			sttp.HandlerValue(
				"ttp.patternId",
				a,
				sttp.GetTtpTmp().SetAnyPatternId,
			)
		}},
		//--- tactic ---
		"ttp.tactic": {func(a any) {
			sttp.HandlerValue(
				"ttp.tactic",
				a,
				sttp.GetTtpTmp().SetAnyTactic,
			)
		}},
		//****************** pattern *******************
		//--- ttp.extraData.pattern.remoteSupport ---
		"ttp.extraData.pattern.remoteSupport": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.remoteSupport",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.extraData.pattern.revoked ---
		"ttp.extraData.pattern.revoked": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.revoked",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRevoked,
			)
		}},
		//--- ttp.extraData.pattern._createdAt ---
		"ttp.extraData.pattern._createdAt": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._createdAt",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.extraData.pattern._createdBy ---
		"ttp.extraData.pattern._createdBy": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._createdBy",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.extraData.pattern._id ---
		"ttp.extraData.pattern._id": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._id",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningId,
			)
		}},
		//--- ttp.extraData.pattern._type ---
		"ttp.extraData.pattern._type": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._type",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningType,
			)
		}},
		//--- ttp.extraData.pattern.detection ---
		"ttp.extraData.pattern.detection": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.detection",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDetection,
			)
		}},
		//--- ttp.extraData.pattern.description ---
		"ttp.extraData.pattern.description": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.description",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDescription,
			)
		}},
		// --- ttp.extraData.pattern.name ---
		"ttp.extraData.pattern.name": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.name",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyName,
			)
		}},
		// --- ttp.extraData.pattern.patternId ---
		"ttp.extraData.pattern.patternId": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.patternId",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternId,
			)
		}},
		// --- ttp.extraData.pattern.patternType ---
		"ttp.extraData.pattern.patternType": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.patternType",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternType,
			)
		}},
		//--- ttp.extraData.pattern.url ---
		"ttp.extraData.pattern.url": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.url",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyURL,
			)
		}},
		//--- ttp.extraData.pattern.version ---
		"ttp.extraData.pattern.version": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.version",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyVersion,
			)
		}},
		//--- ttp.extraData.pattern.platforms ---
		"ttp.extraData.pattern.platforms": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.platforms",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPlatforms,
			)
		}},
		//--- ttp.extraData.pattern.permissionsRequired ---
		"ttp.extraData.pattern.permissionsRequired": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.permissionsRequired",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.extraData.pattern.dataSources ---
		"ttp.extraData.pattern.dataSources": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.dataSources",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDataSources,
			)
		}},
		//--- ttp.extraData.pattern.tactics ---
		"ttp.extraData.pattern.tactics": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.tactics",
				a,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyTactics,
			)
		}},
		//****************** patternParent *******************
		//--- ttp.extraData.patternParent.remoteSupport ---
		"ttp.extraData.patternParent.remoteSupport": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.remoteSupport",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.extraData.patternParent.revoked ---
		"ttp.extraData.patternParent.revoked": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.revoked",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRevoked,
			)
		}},
		//--- ttp.extraData.patternParent._createdAt ---
		"ttp.extraData.patternParent._createdAt": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._createdAt",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.extraData.patternParent._createdBy ---
		"ttp.extraData.patternParent._createdBy": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._createdBy",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.extraData.patternParent._id ---
		"ttp.extraData.patternParent._id": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._id",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningId,
			)
		}},
		//--- ttp.extraData.patternParent._type ---
		"ttp.extraData.patternParent._type": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._type",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningType,
			)
		}},
		//--- ttp.extraData.patternParent.detection ---
		"ttp.extraData.patternParent.detection": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.detection",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDetection,
			)
		}},
		//--- ttp.extraData.patternParent.description ---
		"ttp.extraData.patternParent.description": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.description",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDescription,
			)
		}},
		// --- ttp.extraData.patternParent.name ---
		"ttp.extraData.patternParent.name": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.name",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyName,
			)
		}},
		// --- ttp.extraData.patternParent.patternId ---
		"ttp.extraData.patternParent.patternId": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.patternId",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternId,
			)
		}},
		// --- ttp.extraData.patternParent.patternType ---
		"ttp.extraData.patternParent.patternType": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.patternType",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternType,
			)
		}},
		//--- ttp.extraData.patternParent.url ---
		"ttp.extraData.patternParent.url": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.url",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyURL,
			)
		}},
		//--- ttp.extraData.patternParent.version ---
		"ttp.extraData.patternParent.version": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.version",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyVersion,
			)
		}},
		//--- ttp.extraData.patternParent.platforms ---
		"ttp.extraData.patternParent.platforms": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.platforms",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPlatforms,
			)
		}},
		//--- ttp.extraData.patternParent.permissionsRequired ---
		"ttp.extraData.patternParent.permissionsRequired": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.permissionsRequired",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.extraData.patternParent.dataSources ---
		"ttp.extraData.patternParent.dataSources": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.dataSources",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDataSources,
			)
		}},
		//--- ttp.extraData.patternParent.tactics ---
		"ttp.extraData.patternParent.tactics": {func(a any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.tactics",
				a,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyTactics,
			)
		}},
	}
}
