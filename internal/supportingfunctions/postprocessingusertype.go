package supportingfunctions

import "strings"

type UserTypeGetter interface {
	GetData() string
	GetDataType() string
	SetValueData(string)
	SetValueSensorId(string)
	SetValueSnortSid(string)
	SetAnySnortSidNumber(any)
	SetValueSnortSidNumber(int)
}

// PostProcessingUserType выполняет постобработку некоторых пользовательских типов
func PostProcessingUserType[T UserTypeGetter](ut T) (T, bool) {
	handlers := map[string]func(utg UserTypeGetter){
		"snort_sid": func(utg UserTypeGetter) {
			if !strings.Contains(utg.GetData(), ",") {
				if utg.GetData() != "" {
					utg.SetValueSnortSid(utg.GetData())
					utg.SetAnySnortSidNumber(utg.GetData())
				}

				return
			}

			tmp := strings.SplitSeq(utg.GetData(), ",")
			for v := range tmp {
				str := strings.TrimSpace(v)
				utg.SetValueSnortSid(str)
				utg.SetAnySnortSidNumber(str)
			}
		},
		"ip_home": func(utg UserTypeGetter) {
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
