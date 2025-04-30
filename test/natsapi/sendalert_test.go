package nats_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/confighandler"
)

func TestSendAlert(t *testing.T) {
	var (
		conf *confighandler.ConfigApp
		nc   *nats.Conn
		b    []byte

		err error
	)

	t.Setenv("GO_PHDOCBASEDB_MAIN", "test")

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	t.Run("Тест 1. Чтение конфигурационного файла", func(t *testing.T) {
		conf, err = confighandler.New("placeholder_doc-base_db")
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Инициализация подключения к NATS", func(t *testing.T) {
		nc, err = nats.Connect(
			fmt.Sprintf("%s:%d", conf.GetNATS().Host, conf.GetNATS().Port),
			nats.MaxReconnects(-1),
			nats.ReconnectWait(3*time.Second))
		assert.NoError(t, err)

		// обработка разрыва соединения с NATS
		nc.SetDisconnectErrHandler(func(c *nats.Conn, err error) {
			log.Println(err)
		})

		// обработка переподключения к NATS
		nc.SetReconnectHandler(func(c *nats.Conn) {
			log.Println(err)
		})
	})

	t.Run("Тест 3. Чтение тестовых данных", func(t *testing.T) {
		b, err = os.ReadFile("../test_json/alert.json")
		assert.NoError(t, err)
	})

	t.Run("Тест 4. Передача данных в NATS", func(t *testing.T) {
		chDone := make(chan struct{})
		confDb := conf.GetStorageDB()

		go func(nc *nats.Conn) {
			nc.Subscribe(confDb.Storage["alert"], func(msg *nats.Msg) {
				fmt.Printf("Received command: %s\n", string(msg.Data))

				chDone <- struct{}{}
			})
		}(nc)

		t.Log("GO_PHDOCBASEDB_MAIN =", os.Getenv("GO_PHDOCBASEDB_MAIN"))
		t.Log("conf.Subscriptions.ListenerCase:", confDb.Storage["alert"])

		time.Sleep(3 * time.Second)

		err = nc.Publish(confDb.Storage["alert"], b)
		assert.NoError(t, err)

		fmt.Println("Before")

		<-chDone

		fmt.Println("After")
	})

	t.Cleanup(func() {
		nc.Close()
	})
}
