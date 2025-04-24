package handlers

// NewListHandlerTtp обработчик событий типа 'ttps.*' для объекта 'case'
func NewListHandlerTtp(sttp *SupportiveTtp) map[string][]func(any) {
	return map[string][]func(any){
		//--- occurDate ---
		"ttp.occurDate": {func(i any) {
			sttp.HandlerValue(
				"ttp.occurDate",
				i,
				sttp.GetTtpTmp().SetAnyOccurDate,
			)
		}},
		//--- _createdAt ---
		"ttp._createdAt": {func(i any) {
			sttp.HandlerValue(
				"ttp._createdAt",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _id ---
		"ttp._id": {func(i any) {
			sttp.HandlerValue(
				"ttp._id",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningId,
			)
		}},
		//--- _createdBy ---
		"ttp._createdBy": {func(i any) {
			sttp.HandlerValue(
				"ttp._createdBy",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- patternId ---
		"ttp.patternId": {func(i any) {
			sttp.HandlerValue(
				"ttp.patternId",
				i,
				sttp.GetTtpTmp().SetAnyPatternId,
			)
		}},
		//--- tactic ---
		"ttp.tactic": {func(i any) {
			sttp.HandlerValue(
				"ttp.tactic",
				i,
				sttp.GetTtpTmp().SetAnyTactic,
			)
		}},
		//****************** pattern *******************
		//--- ttp.extraData.pattern.remoteSupport ---
		"ttp.extraData.pattern.remoteSupport": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.remoteSupport",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.extraData.pattern.revoked ---
		"ttp.extraData.pattern.revoked": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.revoked",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRevoked,
			)
		}},
		//--- ttp.extraData.pattern._createdAt ---
		"ttp.extraData.pattern._createdAt": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._createdAt",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.extraData.pattern._createdBy ---
		"ttp.extraData.pattern._createdBy": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._createdBy",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.extraData.pattern._id ---
		"ttp.extraData.pattern._id": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._id",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningId,
			)
		}},
		//--- ttp.extraData.pattern._type ---
		"ttp.extraData.pattern._type": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern._type",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningType,
			)
		}},
		//--- ttp.extraData.pattern.detection ---
		"ttp.extraData.pattern.detection": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.detection",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDetection,
			)
		}},
		//--- ttp.extraData.pattern.description ---
		"ttp.extraData.pattern.description": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.description",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDescription,
			)
		}},
		// --- ttp.extraData.pattern.name ---
		"ttp.extraData.pattern.name": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.name",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyName,
			)
		}},
		// --- ttp.extraData.pattern.patternId ---
		"ttp.extraData.pattern.patternId": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.patternId",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternId,
			)
		}},
		// --- ttp.extraData.pattern.patternType ---
		"ttp.extraData.pattern.patternType": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.patternType",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternType,
			)
		}},
		//--- ttp.extraData.pattern.url ---
		"ttp.extraData.pattern.url": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.url",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyURL,
			)
		}},
		//--- ttp.extraData.pattern.version ---
		"ttp.extraData.pattern.version": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.version",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyVersion,
			)
		}},
		//--- ttp.extraData.pattern.platforms ---
		"ttp.extraData.pattern.platforms": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.platforms",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPlatforms,
			)
		}},
		//--- ttp.extraData.pattern.permissionsRequired ---
		"ttp.extraData.pattern.permissionsRequired": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.permissionsRequired",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.extraData.pattern.dataSources ---
		"ttp.extraData.pattern.dataSources": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.dataSources",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDataSources,
			)
		}},
		//--- ttp.extraData.pattern.tactics ---
		"ttp.extraData.pattern.tactics": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.pattern.tactics",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyTactics,
			)
		}},
		//****************** patternParent *******************
		//--- ttp.extraData.patternParent.remoteSupport ---
		"ttp.extraData.patternParent.remoteSupport": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.remoteSupport",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.extraData.patternParent.revoked ---
		"ttp.extraData.patternParent.revoked": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.revoked",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRevoked,
			)
		}},
		//--- ttp.extraData.patternParent._createdAt ---
		"ttp.extraData.patternParent._createdAt": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._createdAt",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.extraData.patternParent._createdBy ---
		"ttp.extraData.patternParent._createdBy": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._createdBy",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.extraData.patternParent._id ---
		"ttp.extraData.patternParent._id": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._id",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningId,
			)
		}},
		//--- ttp.extraData.patternParent._type ---
		"ttp.extraData.patternParent._type": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent._type",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningType,
			)
		}},
		//--- ttp.extraData.patternParent.detection ---
		"ttp.extraData.patternParent.detection": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.detection",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDetection,
			)
		}},
		//--- ttp.extraData.patternParent.description ---
		"ttp.extraData.patternParent.description": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.description",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDescription,
			)
		}},
		// --- ttp.extraData.patternParent.name ---
		"ttp.extraData.patternParent.name": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.name",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyName,
			)
		}},
		// --- ttp.extraData.patternParent.patternId ---
		"ttp.extraData.patternParent.patternId": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.patternId",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternId,
			)
		}},
		// --- ttp.extraData.patternParent.patternType ---
		"ttp.extraData.patternParent.patternType": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.patternType",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternType,
			)
		}},
		//--- ttp.extraData.patternParent.url ---
		"ttp.extraData.patternParent.url": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.url",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyURL,
			)
		}},
		//--- ttp.extraData.patternParent.version ---
		"ttp.extraData.patternParent.version": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.version",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyVersion,
			)
		}},
		//--- ttp.extraData.patternParent.platforms ---
		"ttp.extraData.patternParent.platforms": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.platforms",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPlatforms,
			)
		}},
		//--- ttp.extraData.patternParent.permissionsRequired ---
		"ttp.extraData.patternParent.permissionsRequired": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.permissionsRequired",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.extraData.patternParent.dataSources ---
		"ttp.extraData.patternParent.dataSources": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.dataSources",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDataSources,
			)
		}},
		//--- ttp.extraData.patternParent.tactics ---
		"ttp.extraData.patternParent.tactics": {func(i any) {
			sttp.HandlerValue(
				"ttp.extraData.patternParent.tactics",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyTactics,
			)
		}},
	}
}
