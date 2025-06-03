package alertgenerator_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

type ChMessage struct {
	Type    string
	Message string
}

func NewChMessage() *ChMessage {
	return &ChMessage{}
}

func (chm *ChMessage) GetType() string {
	return chm.Type
}

func (chm *ChMessage) GetMessage() string {
	return chm.Message
}

func (chm *ChMessage) SetType(v string) {
	chm.Type = v
}

func (chm *ChMessage) SetMessage(v string) {
	chm.Message = v
}

type Logging struct {
	ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{ch: make(chan interfaces.Messager)}
}

func (l *Logging) SetChan(v chan interfaces.Messager) {
	l.ch = v
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.ch
}

func (l *Logging) Send(msgType, msgData string) {
	msg := NewChMessage()
	msg.SetType(msgType)
	msg.SetMessage(msgData)

	l.ch <- msg
}

func TestCaseGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	logging := NewLogging()
	counting := countermessage.New(logging.ch)
	counting.Start(ctx)

	b, err := os.ReadFile("../test_json/alert.json")
	assert.NoError(t, err)

	decoder := decoderjsondocuments.New(counting, logging)
	chDecode := decoder.Start(b, "taskId_628292h")

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-logging.GetChan():
				fmt.Println("Log:", msg)

			}
		}
	}()

	rootId, verifyCase, listRawFields := documentgenerator.AlertGenerator(chDecode)
	assert.Equal(t, rootId, "~91686760480")

	bjson, err := json.MarshalIndent(verifyCase, "", " ")
	assert.NoError(t, err)

	fmt.Println(string(bjson))
	fmt.Println("--------------------")
	fmt.Println("listRawFields:", listRawFields)

	cancel()
}
