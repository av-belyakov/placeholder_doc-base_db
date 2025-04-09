package decoderjsondocuments

import (
	"encoding/json"
	"errors"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// New конструктор нового обработчика JSON сообщения
func New(counter *countermessage.CounterMessage, logger interfaces.Logger) *DecoderJsonDocuments {
	return &DecoderJsonDocuments{logger: logger, counter: counter}
}

// Start инициализация обработки
func (s *DecoderJsonDocuments) Start(b []byte, taskId string) chan interfaces.CustomJsonDecoder {
	chanInput := make(chan interfaces.CustomJsonDecoder)

	go func() {
		//для карт
		listMap := map[string]any{}
		if err := json.Unmarshal(b, &listMap); err == nil {
			if len(listMap) == 0 {
				s.logger.Send("error", supportingfunctions.CustomError(errors.New("error decoding the json message, it may be empty")).Error())

				return
			}

			_ = processingReflectMap(chanInput, listMap, "")
		} else {
			// для срезов
			listSlice := []any{}
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

		// счетчик обработанных кейсов
		s.counter.SendMessage("update processed events", 1)

		close(chanInput)
	}()

	return chanInput
}
