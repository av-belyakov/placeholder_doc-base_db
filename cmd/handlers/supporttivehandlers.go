package handlers

import (
	"strings"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// PostProcessingUserType выполняет постобработку некоторых пользовательских типов
func PostProcessingUserType[T interfaces.UserTypeGetter](ut T) (T, bool) {
	handlers := map[string]func(utg interfaces.UserTypeGetter){
		"snort_sid": func(utg interfaces.UserTypeGetter) {
			if !strings.Contains(utg.GetData(), ",") {
				if utg.GetData() != "" {
					utg.SetValueSnortSid(utg.GetData())
					utg.SetAnySnortSidNumber(utg.GetData())
				}

				return
			}

			tmp := strings.Split(utg.GetData(), ",")
			for _, v := range tmp {
				str := strings.TrimSpace(v)
				utg.SetValueSnortSid(str)
				utg.SetAnySnortSidNumber(str)
			}
		},
		"ip_home": func(utg interfaces.UserTypeGetter) {
			if !strings.Contains(utg.GetData(), ":") {
				return
			}

			tmp := strings.Split(utg.GetData(), ":")
			if len(tmp) != 2 {
				return
			}

			utg.SetValueSensorId(tmp[0])
			utg.SetValueData(tmp[1])
		},
	}

	f, ok := handlers[ut.GetDataType()]
	if !ok {
		return ut, false
	}

	f(ut)

	return ut, true
}
