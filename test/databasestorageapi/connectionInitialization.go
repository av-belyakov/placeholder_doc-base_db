package databasestorageapi_test

import (
	"context"
	"fmt"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// *** для счётчика ***
type Counting struct{}

func (c *Counting) SendMessage(msg string, count int) {}

// *** для сообщений ***
type Message struct {
	Type, Message string
}

func (m *Message) GetType() string {
	return m.Type
}

func (m *Message) SetType(v string) {
	m.Type = v
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetMessage(v string) {
	m.Message = v
}

// *** для логирования ***
type Logging struct {
	ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{
		ch: make(chan interfaces.Messager),
	}
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.ch
}

func (l *Logging) Send(msgType, msgData string) {
	l.ch <- &Message{Type: msgType, Message: msgData}
}

func ConnectionInitialization(ctx context.Context, opts ...databasestorageapi.DatabaseStorageOptions) (*databasestorageapi.DatabaseStorage, error) {
	var (
		apiDBS *databasestorageapi.DatabaseStorage
		err    error
	)

	logging := NewLogging()
	counting := &Counting{}

	go func(ctx context.Context, l *Logging) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

			}
		}
	}(ctx, logging)

	apiDBS, err = databasestorageapi.New(counting, logging, opts...)
	if err != nil {
		return apiDBS, err
	}

	if err := apiDBS.Start(ctx); err != nil {
		return apiDBS, err
	}

	return apiDBS, err
}
