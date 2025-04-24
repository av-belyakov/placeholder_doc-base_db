package handlers

import (
	"fmt"
	"strings"

	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewListHandlerEventObjectCustomFields обработчик событий 'event.object.customField.*' типа для объектов 'alert' или 'case'
func NewListHandlerEventObjectCustomFields(customFields commonhf.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(any){
		//------------- для обработки тегов содержащих geoip -------------
		"event.object.tags": {func(i any) {
			s := fmt.Sprint(i)
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
		"event.object.customFields.attack-type.order": {func(i any) {
			//создаем элемент "attack-type" если его нет
			NewCustomFieldsElement("attack-type", "string", &customFields)
			_, _, _, str := customFields["attack-type"].Get()
			customFields["attack-type"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.attack-type.string": {func(i any) {
			NewCustomFieldsElement("attack-type", "string", &customFields)
			_, order, _, _ := customFields["attack-type"].Get()
			customFields["attack-type"].Set(order, i)
		}},
		//------------------ class-attack ------------------
		"event.object.customFields.class-attack.order": {func(i any) {
			NewCustomFieldsElement("class-attack", "string", &customFields)
			_, _, _, str := customFields["class-attack"].Get()
			customFields["class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-attack.string": {func(i any) {
			NewCustomFieldsElement("class-attack", "string", &customFields)
			_, order, _, _ := customFields["class-attack"].Get()
			customFields["class-attack"].Set(order, i)
		}},
		//------------------ class-ca ------------------
		"event.object.customFields.class-ca.order": {func(i any) {
			NewCustomFieldsElement("class-ca", "string", &customFields)
			_, _, _, str := customFields["class-ca"].Get()
			customFields["class-ca"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-ca.string": {func(i any) {
			NewCustomFieldsElement("class-ca", "string", &customFields)
			_, order, _, _ := customFields["class-ca"].Get()
			customFields["class-ca"].Set(order, i)
		}},
		//--------------- count-of-files ------------------
		"event.object.customFields.count-of-files.order": {func(i any) {
			NewCustomFieldsElement("count-of-files", "integer", &customFields)
			_, _, _, str := customFields["count-of-files"].Get()
			customFields["count-of-files"].Set(i, str)
		}},
		"event.object.customFields.count-of-files.integer": {func(i any) {
			NewCustomFieldsElement("count-of-files", "integer", &customFields)
			_, order, _, _ := customFields["count-of-files"].Get()
			customFields["count-of-files"].Set(order, i)
		}},
		//--------------- count-of-malwares ------------------
		"event.object.customFields.count-of-malwares.order": {func(i any) {
			NewCustomFieldsElement("count-of-malwares", "integer", &customFields)
			_, _, _, str := customFields["count-of-malwares"].Get()
			customFields["count-of-malwares"].Set(i, str)
		}},
		"event.object.customFields.count-of-malwares.integer": {func(i any) {
			NewCustomFieldsElement("count-of-malwares", "integer", &customFields)
			_, order, _, _ := customFields["count-of-malwares"].Get()
			customFields["count-of-malwares"].Set(order, i)
		}},
		//--------------- event-number ------------------
		"event.object.customFields.event-number.order": {func(i any) {
			NewCustomFieldsElement("event-number", "integer", &customFields)
			_, _, _, str := customFields["event-number"].Get()
			customFields["event-number"].Set(i, str)
		}},
		"event.object.customFields.event-number.integer": {func(i any) {
			NewCustomFieldsElement("event-number", "integer", &customFields)
			_, order, _, _ := customFields["event-number"].Get()
			customFields["event-number"].Set(order, i)
		}},
		//--------------- external-letter ------------------
		"event.object.customFields.external-letter.order": {func(i any) {
			NewCustomFieldsElement("external-letter", "integer", &customFields)
			_, _, _, str := customFields["external-letter"].Get()
			customFields["external-letter"].Set(i, str)
		}},
		"event.object.customFields.external-letter.integer": {func(i any) {
			NewCustomFieldsElement("external-letter", "integer", &customFields)
			_, order, _, _ := customFields["external-letter"].Get()
			customFields["external-letter"].Set(order, i)
		}},
		//--------------- misp-event-id ------------------
		"event.object.customFields.misp-event-id.order": {func(i any) {
			NewCustomFieldsElement("misp-event-id", "string", &customFields)
			_, _, _, str := customFields["misp-event-id"].Get()
			customFields["misp-event-id"].Set(i, str)
		}},
		"event.object.customFields.misp-event-id.string": {func(i any) {
			NewCustomFieldsElement("misp-event-id", "string", &customFields)
			_, order, _, _ := customFields["misp-event-id"].Get()
			customFields["misp-event-id"].Set(order, i)
		}},
		// --------------- verdict ------------------
		"event.object.customFields.verdict.order": {func(i any) {
			NewCustomFieldsElement("verdict", "string", &customFields)
			_, _, _, str := customFields["verdict"].Get()
			customFields["verdict"].Set(i, str)
		}},
		"event.object.customFields.verdict.string": {func(i any) {
			NewCustomFieldsElement("verdict", "string", &customFields)
			_, order, _, _ := customFields["verdict"].Get()
			customFields["verdict"].Set(order, i)
		}},
		// --------------- classification ------------------
		"event.object.customFields.classification.order": {func(i any) {
			NewCustomFieldsElement("classification", "string", &customFields)
			_, _, _, str := customFields["classification"].Get()
			customFields["classification"].Set(i, str)
		}},
		"event.object.customFields.classification.string": {func(i any) {
			NewCustomFieldsElement("classification", "string", &customFields)
			_, order, _, _ := customFields["classification"].Get()
			customFields["classification"].Set(order, i)
		}},
		//--------------- gratitude ------------------ номер благодарственного письма ????
		"event.object.customFields.gratitude.order": {func(i any) {
			NewCustomFieldsElement("gratitude", "integer", &customFields)
			_, _, _, str := customFields["gratitude"].Get()
			customFields["gratitude"].Set(i, str)
		}},
		"event.object.customFields.gratitude.integer": {func(i any) {
			NewCustomFieldsElement("gratitude", "integer", &customFields)
			_, order, _, _ := customFields["gratitude"].Get()
			customFields["gratitude"].Set(order, i)
		}},
		//------------------ ncircc-class-attack ------------------
		"event.object.customFields.ncircc-class-attack.order": {func(i any) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &customFields)
			_, _, _, str := customFields["ncircc-class-attack"].Get()
			customFields["ncircc-class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ncircc-class-attack.string": {func(i any) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &customFields)
			_, order, _, _ := customFields["ncircc-class-attack"].Get()
			customFields["ncircc-class-attack"].Set(order, i)
		}},
		//------------------ inbox1 ------------------
		"event.object.customFields.inbox1.order": {func(i any) {
			NewCustomFieldsElement("inbox1", "string", &customFields)
			_, _, _, str := customFields["inbox1"].Get()
			customFields["inbox1"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inbox1.string": {func(i any) {
			NewCustomFieldsElement("inbox1", "string", &customFields)
			_, order, _, _ := customFields["inbox1"].Get()
			customFields["inbox1"].Set(order, i)
		}},
		//------------------ inner-letter ------------------
		"event.object.customFields.inner-letter.order": {func(i any) {
			NewCustomFieldsElement("inner-letter", "string", &customFields)
			_, _, _, str := customFields["inner-letter"].Get()
			customFields["inner-letter"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inner-letter.string": {func(i any) {
			NewCustomFieldsElement("inner-letter", "string", &customFields)
			_, order, _, _ := customFields["inner-letter"].Get()
			customFields["inner-letter"].Set(order, i)
		}},
		//------------------ notification ------------------
		"event.object.customFields.notification.order": {func(i any) {
			NewCustomFieldsElement("notification", "string", &customFields)
			_, _, _, str := customFields["notification"].Get()
			customFields["notification"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ report ------------------
		"event.object.customFields.report.order": {func(i any) {
			NewCustomFieldsElement("report", "string", &customFields)
			_, _, _, str := customFields["report"].Get()
			customFields["report"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ first-time ------------------
		"event.object.customFields.first-time.order": {func(i any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields["first-time"].Get()
			customFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields["first-time"].Get()
			customFields["first-time"].Set(order, i)
		}},
		//------------------ last-time ------------------
		"event.object.customFields.last-time.order": {func(i any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields["last-time"].Get()
			customFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields["last-time"].Get()
			customFields["last-time"].Set(order, i)
		}},
		//------------------ sphere ------------------
		"event.object.customFields.sphere.order": {func(i any) {
			NewCustomFieldsElement("sphere", "string", &customFields)
			_, _, _, str := customFields["sphere"].Get()
			customFields["sphere"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.sphere.string": {func(i any) {
			NewCustomFieldsElement("sphere", "string", &customFields)
			_, order, _, _ := customFields["sphere"].Get()
			customFields["sphere"].Set(order, i)
		}},
		//------------------ state ------------------
		"event.object.customFields.state.order": {func(i any) {
			NewCustomFieldsElement("state", "string", &customFields)
			_, _, _, str := customFields["state"].Get()
			customFields["state"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.state.string": {func(i any) {
			NewCustomFieldsElement("state", "string", &customFields)
			_, order, _, _ := customFields["state"].Get()
			customFields["state"].Set(order, i)
		}},
		//------------------ ir-name ------------------
		"event.object.customFields.ir-name.order": {func(i any) {
			NewCustomFieldsElement("ir-name", "string", &customFields)
			_, _, _, str := customFields["ir-name"].Get()
			customFields["ir-name"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ir-name.string": {func(i any) {
			NewCustomFieldsElement("ir-name", "string", &customFields)
			_, order, _, _ := customFields["ir-name"].Get()
			customFields["ir-name"].Set(order, i)
		}},
		//------------------ id-soa ------------------
		"event.object.customFields.id-soa.order": {func(i any) {
			NewCustomFieldsElement("id-soa", "string", &customFields)
			_, _, _, str := customFields["id-soa"].Get()
			customFields["id-soa"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.id-soa.string": {func(i any) {
			NewCustomFieldsElement("id-soa", "string", &customFields)
			_, order, _, _ := customFields["id-soa"].Get()
			customFields["id-soa"].Set(order, i)
		}},
		//--------------- is-incident ------------------
		"event.object.customFields.is-incident.order": {func(i any) {
			NewCustomFieldsElement("is-incident", "boolen", &customFields)
			_, _, _, str := customFields["is-incident"].Get()
			customFields["is-incident"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.is-incident.boolean": {func(i any) {
			NewCustomFieldsElement("is-incident", "boolean", &customFields)
			_, order, _, _ := customFields["is-incident"].Get()
			customFields["is-incident"].Set(order, i)
		}},
		//--------------- work-admin ------------------
		"event.object.customFields.work-admin.order": {func(i any) {
			NewCustomFieldsElement("work-admin", "boolen", &customFields)
			_, _, _, str := customFields["work-admin"].Get()
			customFields["work-admin"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.work-admin.boolean": {func(i any) {
			NewCustomFieldsElement("work-admin", "boolean", &customFields)
			_, order, _, _ := customFields["work-admin"].Get()
			customFields["work-admin"].Set(order, i)
		}},
	}
}
