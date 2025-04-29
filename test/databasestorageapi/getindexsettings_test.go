package databasestorageapi

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func TestGetIndexSettings(t *testing.T) {
	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	logging := NewLogging()

	counting := &Counting{}

	go func(ctx context.Context, l *Logging) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-l.GetChan():
				t.Log("LOG:", msg)

			}
		}
	}(ctx, logging)

	apiDBS, err := databasestorageapi.New(
		counting,
		logging,
		databasestorageapi.WithHost("datahook.cloud.gcm"),
		databasestorageapi.WithPort(9200),
		databasestorageapi.WithUser("writer"),
		databasestorageapi.WithPasswd(os.Getenv("GO_PHDOCBASEDB_DBWLOGPASSWD")),
		databasestorageapi.WithStorage(map[string]string{
			"alert": "module_placeholderdb_alert",
			"case":  "module_placeholderdb_case",
		}))
	if err != nil {
		log.Fatalln(err)
	}
	apiDBS.Start(ctx)

	indexSettings, err := apiDBS.GetIndexSetting(ctx, "module_placeholder_new_case_2025_4", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, len(indexSettings))

	for k, v := range indexSettings {
		t.Logf("Index:'%s'\n\tValue:'%+v'\n", k, v)
	}

	cancel()
}
