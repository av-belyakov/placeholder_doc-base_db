package handlers

import casedetailshf "github.com/av-belyakov/objectsthehiveformat/casedetails"

// NewListHandlerEventCaseDetails обработчик событий типа 'case' для 'event.details'
func NewListHandlerEventCaseDetails(details *casedetailshf.EventCaseDetails) map[string][]func(interface{}) {
	return map[string][]func(any){
		"event.details.endDate":          {details.SetAnyEndDate},
		"event.details.resolutionStatus": {details.SetAnyResolutionStatus},
		"event.details.summary":          {details.SetAnySummary},
		"event.details.status":           {details.SetAnyStatus},
		"event.details.impactStatus":     {details.SetAnyImpactStatus},
	}
}
