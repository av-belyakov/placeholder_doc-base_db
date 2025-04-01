// Пакет countermassage подсчитывает входящие сообщения и отправляет их системе мониторинга
package countermessage

import (
	"context"
	"fmt"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/informationcountingstorage"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// New конструктор счетчика сообщений
func New(chOutput chan<- interfaces.Messager) *CounterMessage {
	return &CounterMessage{
		storage:  informationcountingstorage.NewInformationMessageCountingStorage(),
		chOutput: chOutput,
		chInput:  make(chan DataCounterMessage),
	}
}

// Start обработчик подсчитывающий входящие данные
func (c *CounterMessage) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(c.chInput)

				return

			case data := <-c.chInput:
				d, h, m, s := supportingfunctions.GetDifference(c.storage.GetStartTime(), time.Now())
				patternTime := fmt.Sprintf("со старта приложения: дней %d, часов %d, минут %d, секунд %d", d, h, m, s)

				var count uint = uint(data.count)
				c.storage.Increase(data.msgType)
				if v, err := c.storage.GetCount(data.msgType); err == nil {
					count = v
				}

				var msg string = data.msgType
				switch data.msgType {
				case "update accepted events":
					msg = fmt.Sprintf("принято: %d, %s", count, patternTime)

				case "update processed events":
					msg = fmt.Sprintf("обработано: %d, %s", count, patternTime)

				case "update events meet rules - subject_case":
					msg = fmt.Sprintf("подписка-'subject_case', соответствие правилам: %d, %s", count, patternTime)

				case "update events meet rules - subject_alert":
					msg = fmt.Sprintf("подписка-'subject_alert', соответствие правилам: %d, %s", count, patternTime)

				case "update count insert subject_case to db":
					msg = fmt.Sprintf("подписка-'subject_case', добавлено в MongoDB: %d, %s", count, patternTime)

				case "update count insert subject_alert to db":
					msg = fmt.Sprintf("подписка-'subject_alert', добавлено в MongoDB: %d, %s", count, patternTime)

				}

				message := NewSomeMessage()
				message.SetType("info")
				message.SetMessage(msg)

				c.chOutput <- message
			}
		}
	}()
}

// SendMessage отправка количественных значений счетчику сообщений
func (c *CounterMessage) SendMessage(msgType string, count int) {
	c.chInput <- DataCounterMessage{msgType: msgType, count: count}
}
