package documentgenerator

import (
	"encoding/json"
	"errors"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func New(logger interfaces.Logger, counter *countermessage.CounterMessage) *HandlerJsonMessageSettings {
	return &HandlerJsonMessageSettings{
		logger:  logger,
		counter: counter,
	}
}

// Start инициализация обработки
func (s *HandlerJsonMessageSettings) Start(b []byte, taskId string) chan ChanInputCreateNewFormat {
	chanInput := make(chan ChanInputCreateNewFormat)

	go func() {
		//для карт
		listMap := map[string]interface{}{}
		if err := json.Unmarshal(b, &listMap); err == nil {
			if len(listMap) == 0 {
				s.logger.Send("error", supportingfunctions.CustomError(errors.New("error decoding the json message, it may be empty")).Error())

				return
			}

			_ = processingReflectMap(chanInput, listMap, "")
		} else {
			// для срезов
			listSlice := []interface{}{}
			if err = json.Unmarshal(b, &listSlice); err != nil {
				s.logger.Send("error", supportingfunctions.CustomError(err).Error())

				return
			}

			if len(listSlice) == 0 {
				s.logger.Send("error", supportingfunctions.CustomError(errors.New("error decoding the json message, it may be empty")).Error())

				return
			}

			_ = processingReflectSlice(chanInput, listSlice, "")
		}

		// сетчик обработанных кейсов
		s.counter.SendMessage("update processed events", 1)

		close(chanInput)
	}()

	return chanInput
}
