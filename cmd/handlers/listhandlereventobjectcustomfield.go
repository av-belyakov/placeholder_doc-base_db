package handlers

import (
	"fmt"
	"strings"

	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewListHandlerEventObjectCustomFields обработчик событий 'event.object.customField.*' типа для объектов 'alert' или 'case'
func NewListHandlerEventObjectCustomFields(customFields commonhf.CustomFields) map[string][]func(any) {
	return map[string][]func(any){
		//------------- для обработки тегов содержащих geoip -------------
		"event.object.tags": {func(a any) {
			s := fmt.Sprint(a)
			if !strings.Contains(s, "geoip") {
				return
			}

			tmp := strings.Split(s, "=")
			if len(tmp) < 2 {
				return
			}

			//создаем элемент "geoip" если его нет
			NewCustomFieldsElement("geoip", "string", &customFields)
			customFields["geoip"].Set(0, supportingfunctions.TrimIsNotLetter(tmp[1]))
		}},

		//------------------ attack-type ------------------
		"event.object.customFields.attack-type.order": {func(a any) {
			//создаем элемент "attack-type" если его нет
			NewCustomFieldsElement("attack-type", "string", &customFields)
			_, _, _, str := customFields["attack-type"].Get()
			customFields["attack-type"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.attack-type.string": {func(a any) {
			NewCustomFieldsElement("attack-type", "string", &customFields)
			_, order, _, _ := customFields["attack-type"].Get()
			customFields["attack-type"].Set(order, a)
		}},
		//------------------ class-attack ------------------
		"event.object.customFields.class-attack.order": {func(a any) {
			NewCustomFieldsElement("class-attack", "string", &customFields)
			_, _, _, str := customFields["class-attack"].Get()
			customFields["class-attack"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-attack.string": {func(a any) {
			NewCustomFieldsElement("class-attack", "string", &customFields)
			_, order, _, _ := customFields["class-attack"].Get()
			customFields["class-attack"].Set(order, a)
		}},
		//------------------ class-ca ------------------
		"event.object.customFields.class-ca.order": {func(a any) {
			NewCustomFieldsElement("class-ca", "string", &customFields)
			_, _, _, str := customFields["class-ca"].Get()
			customFields["class-ca"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-ca.string": {func(a any) {
			NewCustomFieldsElement("class-ca", "string", &customFields)
			_, order, _, _ := customFields["class-ca"].Get()
			customFields["class-ca"].Set(order, a)
		}},
		//--------------- count-of-files ------------------
		"event.object.customFields.count-of-files.order": {func(a any) {
			NewCustomFieldsElement("count-of-files", "integer", &customFields)
			_, _, _, str := customFields["count-of-files"].Get()
			customFields["count-of-files"].Set(a, str)
		}},
		"event.object.customFields.count-of-files.integer": {func(a any) {
			NewCustomFieldsElement("count-of-files", "integer", &customFields)
			_, order, _, _ := customFields["count-of-files"].Get()
			customFields["count-of-files"].Set(order, a)
		}},
		//--------------- count-of-malwares ------------------
		"event.object.customFields.count-of-malwares.order": {func(a any) {
			NewCustomFieldsElement("count-of-malwares", "integer", &customFields)
			_, _, _, str := customFields["count-of-malwares"].Get()
			customFields["count-of-malwares"].Set(a, str)
		}},
		"event.object.customFields.count-of-malwares.integer": {func(a any) {
			NewCustomFieldsElement("count-of-malwares", "integer", &customFields)
			_, order, _, _ := customFields["count-of-malwares"].Get()
			customFields["count-of-malwares"].Set(order, a)
		}},
		//--------------- event-number ------------------
		"event.object.customFields.event-number.order": {func(a any) {
			NewCustomFieldsElement("event-number", "integer", &customFields)
			_, _, _, str := customFields["event-number"].Get()
			customFields["event-number"].Set(a, str)
		}},
		"event.object.customFields.event-number.integer": {func(a any) {
			NewCustomFieldsElement("event-number", "integer", &customFields)
			_, order, _, _ := customFields["event-number"].Get()
			customFields["event-number"].Set(order, a)
		}},
		//--------------- external-letter ------------------
		"event.object.customFields.external-letter.order": {func(a any) {
			NewCustomFieldsElement("external-letter", "integer", &customFields)
			_, _, _, str := customFields["external-letter"].Get()
			customFields["external-letter"].Set(a, str)
		}},
		"event.object.customFields.external-letter.integer": {func(a any) {
			NewCustomFieldsElement("external-letter", "integer", &customFields)
			_, order, _, _ := customFields["external-letter"].Get()
			customFields["external-letter"].Set(order, a)
		}},
		//--------------- misp-event-id ------------------
		"event.object.customFields.misp-event-id.order": {func(a any) {
			NewCustomFieldsElement("misp-event-id", "string", &customFields)
			_, _, _, str := customFields["misp-event-id"].Get()
			customFields["misp-event-id"].Set(a, str)
		}},
		"event.object.customFields.misp-event-id.string": {func(a any) {
			NewCustomFieldsElement("misp-event-id", "string", &customFields)
			_, order, _, _ := customFields["misp-event-id"].Get()
			customFields["misp-event-id"].Set(order, a)
		}},
		// --------------- verdict ------------------
		"event.object.customFields.verdict.order": {func(a any) {
			NewCustomFieldsElement("verdict", "string", &customFields)
			_, _, _, str := customFields["verdict"].Get()
			customFields["verdict"].Set(a, str)
		}},
		"event.object.customFields.verdict.string": {func(a any) {
			NewCustomFieldsElement("verdict", "string", &customFields)
			_, order, _, _ := customFields["verdict"].Get()
			customFields["verdict"].Set(order, a)
		}},
		// --------------- classification ------------------
		"event.object.customFields.classification.order": {func(a any) {
			NewCustomFieldsElement("classification", "string", &customFields)
			_, _, _, str := customFields["classification"].Get()
			customFields["classification"].Set(a, str)
		}},
		"event.object.customFields.classification.string": {func(a any) {
			NewCustomFieldsElement("classification", "string", &customFields)
			_, order, _, _ := customFields["classification"].Get()
			customFields["classification"].Set(order, a)
		}},
		//--------------- gratitude ------------------ номер благодарственного письма ????
		"event.object.customFields.gratitude.order": {func(a any) {
			NewCustomFieldsElement("gratitude", "integer", &customFields)
			_, _, _, str := customFields["gratitude"].Get()
			customFields["gratitude"].Set(a, str)
		}},
		"event.object.customFields.gratitude.integer": {func(a any) {
			NewCustomFieldsElement("gratitude", "integer", &customFields)
			_, order, _, _ := customFields["gratitude"].Get()
			customFields["gratitude"].Set(order, a)
		}},
		//------------------ ncircc-class-attack ------------------
		"event.object.customFields.ncircc-class-attack.order": {func(a any) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &customFields)
			_, _, _, str := customFields["ncircc-class-attack"].Get()
			customFields["ncircc-class-attack"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ncircc-class-attack.string": {func(a any) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &customFields)
			_, order, _, _ := customFields["ncircc-class-attack"].Get()
			customFields["ncircc-class-attack"].Set(order, a)
		}},
		//------------------ inbox1 ------------------
		"event.object.customFields.inbox1.order": {func(a any) {
			NewCustomFieldsElement("inbox1", "string", &customFields)
			_, _, _, str := customFields["inbox1"].Get()
			customFields["inbox1"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inbox1.string": {func(a any) {
			NewCustomFieldsElement("inbox1", "string", &customFields)
			_, order, _, _ := customFields["inbox1"].Get()
			customFields["inbox1"].Set(order, a)
		}},
		//------------------ inner-letter ------------------
		"event.object.customFields.inner-letter.order": {func(a any) {
			NewCustomFieldsElement("inner-letter", "string", &customFields)
			_, _, _, str := customFields["inner-letter"].Get()
			customFields["inner-letter"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inner-letter.string": {func(a any) {
			NewCustomFieldsElement("inner-letter", "string", &customFields)
			_, order, _, _ := customFields["inner-letter"].Get()
			customFields["inner-letter"].Set(order, a)
		}},
		//------------------ notification ------------------
		"event.object.customFields.notification.order": {func(a any) {
			NewCustomFieldsElement("notification", "string", &customFields)
			_, _, _, str := customFields["notification"].Get()
			customFields["notification"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ report ------------------
		"event.object.customFields.report.order": {func(a any) {
			NewCustomFieldsElement("report", "string", &customFields)
			_, _, _, str := customFields["report"].Get()
			customFields["report"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ first-time ------------------
		"event.object.customFields.first-time.order": {func(a any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields["first-time"].Get()
			customFields["first-time"].Set(a, str)
		}},
		"event.object.customFields.first-time.date": {func(a any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields["first-time"].Get()
			customFields["first-time"].Set(order, a)
		}},
		//------------------ last-time ------------------
		"event.object.customFields.last-time.order": {func(a any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields["last-time"].Get()
			customFields["last-time"].Set(a, str)
		}},
		"event.object.customFields.last-time.date": {func(a any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields["last-time"].Get()
			customFields["last-time"].Set(order, a)
		}},
		//------------------ sphere ------------------
		"event.object.customFields.sphere.order": {func(a any) {
			NewCustomFieldsElement("sphere", "string", &customFields)
			_, _, _, str := customFields["sphere"].Get()
			customFields["sphere"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.sphere.string": {func(a any) {
			NewCustomFieldsElement("sphere", "string", &customFields)
			_, order, _, _ := customFields["sphere"].Get()
			customFields["sphere"].Set(order, a)
		}},
		//------------------ state ------------------
		"event.object.customFields.state.order": {func(a any) {
			NewCustomFieldsElement("state", "string", &customFields)
			_, _, _, str := customFields["state"].Get()
			customFields["state"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.state.string": {func(a any) {
			NewCustomFieldsElement("state", "string", &customFields)
			_, order, _, _ := customFields["state"].Get()
			customFields["state"].Set(order, a)
		}},
		//------------------ ir-name ------------------
		"event.object.customFields.ir-name.order": {func(a any) {
			NewCustomFieldsElement("ir-name", "string", &customFields)
			_, _, _, str := customFields["ir-name"].Get()
			customFields["ir-name"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ir-name.string": {func(a any) {
			NewCustomFieldsElement("ir-name", "string", &customFields)
			_, order, _, _ := customFields["ir-name"].Get()
			customFields["ir-name"].Set(order, a)
		}},
		//------------------ id-soa ------------------
		"event.object.customFields.id-soa.order": {func(a any) {
			NewCustomFieldsElement("id-soa", "string", &customFields)
			_, _, _, str := customFields["id-soa"].Get()
			customFields["id-soa"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.id-soa.string": {func(a any) {
			NewCustomFieldsElement("id-soa", "string", &customFields)
			_, order, _, _ := customFields["id-soa"].Get()
			customFields["id-soa"].Set(order, a)
		}},
		//--------------- is-incident ------------------
		"event.object.customFields.is-incident.order": {func(a any) {
			NewCustomFieldsElement("is-incident", "boolen", &customFields)
			_, _, _, str := customFields["is-incident"].Get()
			customFields["is-incident"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.is-incident.boolean": {func(a any) {
			NewCustomFieldsElement("is-incident", "boolean", &customFields)
			_, order, _, _ := customFields["is-incident"].Get()
			customFields["is-incident"].Set(order, a)
		}},
		//--------------- work-admin ------------------
		"event.object.customFields.work-admin.order": {func(a any) {
			NewCustomFieldsElement("work-admin", "boolen", &customFields)
			_, _, _, str := customFields["work-admin"].Get()
			customFields["work-admin"].Set(a, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.work-admin.boolean": {func(a any) {
			NewCustomFieldsElement("work-admin", "boolean", &customFields)
			_, order, _, _ := customFields["work-admin"].Get()
			customFields["work-admin"].Set(order, a)
		}},
	}
}
