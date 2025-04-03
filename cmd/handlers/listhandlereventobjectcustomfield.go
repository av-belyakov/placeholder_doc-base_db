package handlers

import (
	"fmt"
	"strings"

	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewListHandlerEventObjectCustomFields обработчик событий типа 'case' для 'event.object.customField'
func NewListHandlerEventObjectCustomFields(eventObjectCustomFields commonhf.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//------------- для обработки тегов содержащих geoip -------------
		"event.object.tags": {func(i interface{}) {
			s := fmt.Sprint(i)
			if !strings.Contains(s, "geoip") {
				return
			}

			tmp := strings.Split(s, "=")
			if len(tmp) < 2 {
				return
			}

			//создаем элемент "geoip" если его нет
			NewCustomFieldsElement("geoip", "string", &eventObjectCustomFields)
			eventObjectCustomFields["geoip"].Set(0, supportingfunctions.TrimIsNotLetter(tmp[1]))
		}},

		//------------------ attack-type ------------------
		"event.object.customFields.attack-type.order": {func(i interface{}) {
			//создаем элемент "attack-type" если его нет
			NewCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["attack-type"].Get()
			eventObjectCustomFields["attack-type"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.attack-type.string": {func(i interface{}) {
			NewCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["attack-type"].Get()
			eventObjectCustomFields["attack-type"].Set(order, i)
		}},
		//------------------ class-attack ------------------
		"event.object.customFields.class-attack.order": {func(i interface{}) {
			NewCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["class-attack"].Get()
			eventObjectCustomFields["class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-attack.string": {func(i interface{}) {
			NewCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["class-attack"].Get()
			eventObjectCustomFields["class-attack"].Set(order, i)
		}},
		//------------------ class-ca ------------------
		"event.object.customFields.class-ca.order": {func(i interface{}) {
			NewCustomFieldsElement("class-ca", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["class-ca"].Get()
			eventObjectCustomFields["class-ca"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-ca.string": {func(i interface{}) {
			NewCustomFieldsElement("class-ca", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["class-ca"].Get()
			eventObjectCustomFields["class-ca"].Set(order, i)
		}},
		//--------------- count-of-files ------------------
		"event.object.customFields.count-of-files.order": {func(i interface{}) {
			NewCustomFieldsElement("count-of-files", "integer", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["count-of-files"].Get()
			eventObjectCustomFields["count-of-files"].Set(i, str)
		}},
		"event.object.customFields.count-of-files.integer": {func(i interface{}) {
			NewCustomFieldsElement("count-of-files", "integer", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["count-of-files"].Get()
			eventObjectCustomFields["count-of-files"].Set(order, i)
		}},
		//--------------- count-of-malwares ------------------
		"event.object.customFields.count-of-malwares.order": {func(i interface{}) {
			NewCustomFieldsElement("count-of-malwares", "integer", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["count-of-malwares"].Get()
			eventObjectCustomFields["count-of-malwares"].Set(i, str)
		}},
		"event.object.customFields.count-of-malwares.integer": {func(i interface{}) {
			NewCustomFieldsElement("count-of-malwares", "integer", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["count-of-malwares"].Get()
			eventObjectCustomFields["count-of-malwares"].Set(order, i)
		}},
		//--------------- event-number ------------------
		"event.object.customFields.event-number.order": {func(i interface{}) {
			NewCustomFieldsElement("event-number", "integer", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["event-number"].Get()
			eventObjectCustomFields["event-number"].Set(i, str)
		}},
		"event.object.customFields.event-number.integer": {func(i interface{}) {
			NewCustomFieldsElement("event-number", "integer", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["event-number"].Get()
			eventObjectCustomFields["event-number"].Set(order, i)
		}},
		//--------------- external-letter ------------------
		"event.object.customFields.external-letter.order": {func(i interface{}) {
			NewCustomFieldsElement("external-letter", "integer", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["external-letter"].Get()
			eventObjectCustomFields["external-letter"].Set(i, str)
		}},
		"event.object.customFields.external-letter.integer": {func(i interface{}) {
			NewCustomFieldsElement("external-letter", "integer", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["external-letter"].Get()
			eventObjectCustomFields["external-letter"].Set(order, i)
		}},
		//--------------- misp-event-id ------------------
		"event.object.customFields.misp-event-id.order": {func(i interface{}) {
			NewCustomFieldsElement("misp-event-id", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["misp-event-id"].Get()
			eventObjectCustomFields["misp-event-id"].Set(i, str)
		}},
		"event.object.customFields.misp-event-id.string": {func(i interface{}) {
			NewCustomFieldsElement("misp-event-id", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["misp-event-id"].Get()
			eventObjectCustomFields["misp-event-id"].Set(order, i)
		}},
		// --------------- verdict ------------------
		"event.object.customFields.verdict.order": {func(i interface{}) {
			NewCustomFieldsElement("verdict", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["verdict"].Get()
			eventObjectCustomFields["verdict"].Set(i, str)
		}},
		"event.object.customFields.verdict.string": {func(i interface{}) {
			NewCustomFieldsElement("verdict", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["verdict"].Get()
			eventObjectCustomFields["verdict"].Set(order, i)
		}},
		// --------------- classification ------------------
		"event.object.customFields.classification.order": {func(i interface{}) {
			NewCustomFieldsElement("classification", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["classification"].Get()
			eventObjectCustomFields["classification"].Set(i, str)
		}},
		"event.object.customFields.classification.string": {func(i interface{}) {
			NewCustomFieldsElement("classification", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["classification"].Get()
			eventObjectCustomFields["classification"].Set(order, i)
		}},
		//--------------- gratitude ------------------ номер благодарственного письма ????
		"event.object.customFields.gratitude.order": {func(i interface{}) {
			NewCustomFieldsElement("gratitude", "integer", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["gratitude"].Get()
			eventObjectCustomFields["gratitude"].Set(i, str)
		}},
		"event.object.customFields.gratitude.integer": {func(i interface{}) {
			NewCustomFieldsElement("gratitude", "integer", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["gratitude"].Get()
			eventObjectCustomFields["gratitude"].Set(order, i)
		}},
		//------------------ ncircc-class-attack ------------------
		"event.object.customFields.ncircc-class-attack.order": {func(i interface{}) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields["ncircc-class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ncircc-class-attack.string": {func(i interface{}) {
			NewCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields["ncircc-class-attack"].Set(order, i)
		}},
		//------------------ inbox1 ------------------
		"event.object.customFields.inbox1.order": {func(i interface{}) {
			NewCustomFieldsElement("inbox1", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["inbox1"].Get()
			eventObjectCustomFields["inbox1"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inbox1.string": {func(i interface{}) {
			NewCustomFieldsElement("inbox1", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["inbox1"].Get()
			eventObjectCustomFields["inbox1"].Set(order, i)
		}},
		//------------------ inner-letter ------------------
		"event.object.customFields.inner-letter.order": {func(i interface{}) {
			NewCustomFieldsElement("inner-letter", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["inner-letter"].Get()
			eventObjectCustomFields["inner-letter"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.inner-letter.string": {func(i interface{}) {
			NewCustomFieldsElement("inner-letter", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["inner-letter"].Get()
			eventObjectCustomFields["inner-letter"].Set(order, i)
		}},
		//------------------ notification ------------------
		"event.object.customFields.notification.order": {func(i interface{}) {
			NewCustomFieldsElement("notification", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["notification"].Get()
			eventObjectCustomFields["notification"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ report ------------------
		"event.object.customFields.report.order": {func(i interface{}) {
			NewCustomFieldsElement("report", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["report"].Get()
			eventObjectCustomFields["report"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ first-time ------------------
		"event.object.customFields.first-time.order": {func(i interface{}) {
			NewCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["first-time"].Get()
			eventObjectCustomFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i interface{}) {
			NewCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["first-time"].Get()
			eventObjectCustomFields["first-time"].Set(order, i)
		}},
		//------------------ last-time ------------------
		"event.object.customFields.last-time.order": {func(i interface{}) {
			NewCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["last-time"].Get()
			eventObjectCustomFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i interface{}) {
			NewCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["last-time"].Get()
			eventObjectCustomFields["last-time"].Set(order, i)
		}},
		//------------------ sphere ------------------
		"event.object.customFields.sphere.order": {func(i interface{}) {
			NewCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["sphere"].Get()
			eventObjectCustomFields["sphere"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.sphere.string": {func(i interface{}) {
			NewCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["sphere"].Get()
			eventObjectCustomFields["sphere"].Set(order, i)
		}},
		//------------------ state ------------------
		"event.object.customFields.state.order": {func(i interface{}) {
			NewCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["state"].Get()
			eventObjectCustomFields["state"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.state.string": {func(i interface{}) {
			NewCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["state"].Get()
			eventObjectCustomFields["state"].Set(order, i)
		}},
		//------------------ ir-name ------------------
		"event.object.customFields.ir-name.order": {func(i interface{}) {
			NewCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["ir-name"].Get()
			eventObjectCustomFields["ir-name"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ir-name.string": {func(i interface{}) {
			NewCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["ir-name"].Get()
			eventObjectCustomFields["ir-name"].Set(order, i)
		}},
		//------------------ id-soa ------------------
		"event.object.customFields.id-soa.order": {func(i interface{}) {
			NewCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["id-soa"].Get()
			eventObjectCustomFields["id-soa"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.id-soa.string": {func(i interface{}) {
			NewCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["id-soa"].Get()
			eventObjectCustomFields["id-soa"].Set(order, i)
		}},
		//--------------- is-incident ------------------
		"event.object.customFields.is-incident.order": {func(i interface{}) {
			NewCustomFieldsElement("is-incident", "boolen", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["is-incident"].Get()
			eventObjectCustomFields["is-incident"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.is-incident.boolean": {func(i interface{}) {
			NewCustomFieldsElement("is-incident", "boolean", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["is-incident"].Get()
			eventObjectCustomFields["is-incident"].Set(order, i)
		}},
		//--------------- work-admin ------------------
		"event.object.customFields.work-admin.order": {func(i interface{}) {
			NewCustomFieldsElement("work-admin", "boolen", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["work-admin"].Get()
			eventObjectCustomFields["work-admin"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.work-admin.boolean": {func(i interface{}) {
			NewCustomFieldsElement("work-admin", "boolean", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["work-admin"].Get()
			eventObjectCustomFields["work-admin"].Set(order, i)
		}},
	}
}
