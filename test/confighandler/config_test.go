package confighandler_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/confighandler"
)

const Root_Dir = "placeholder_doc-base_db"

var (
	conf *confighandler.ConfigApp

	err error
)

func TestMain(m *testing.M) {
	os.Unsetenv("GO_PHDOCBASEDB_MAIN")

	//настройки NATS
	os.Unsetenv("GO_PHDOCBASEDB_NHOST")
	os.Unsetenv("GO_PHDOCBASEDB_NPORT")
	os.Unsetenv("GO_PHDOCBASEDB_NCACHETTL")
	os.Unsetenv("GO_PHDOCBASEDB_NSUBSENDERCASE")
	os.Unsetenv("GO_PHDOCBASEDB_NSUBSENDERALERT")
	os.Unsetenv("GO_PHDOCBASEDB_NSUBLISTENERCOMMAND")

	// Настройки доступа к БД в которую будут записыватся alert и case
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEHOST")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPORT")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGENAME")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEUSER")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGENALERT")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGENCASE")

	//настройки доступа к БД в которую будут записыватся логи
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGHOST")
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPORT")
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGNAME")
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGUSER")
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
	os.Unsetenv("GO_PHDOCBASEDB_DBWLOGSTORAGENAME")

	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	os.Setenv("GO_PHDOCBASEDB_MAIN", "test")

	conf, err = confighandler.New(Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}

func TestConfigHandler(t *testing.T) {
	t.Run("Тест чтения конфигурационного файла", func(t *testing.T) {
		t.Run("Тест 1. Проверка настройки NATS из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetNATS().Host, "192.168.9.208")
			assert.Equal(t, conf.GetNATS().Port, 4222)
			assert.Equal(t, conf.GetNATS().CacheTTL, 3600)
			assert.Equal(t, conf.GetNATS().Subscriptions.ListenerAlert, "object.alerttype.test")
			assert.Equal(t, conf.GetNATS().Subscriptions.ListenerCase, "object.casetype.test")
			assert.Equal(t, conf.GetNATS().Subscriptions.SenderCommand, "object.commandstype.test")
		})

		t.Run("Тест 2. Проверка настройки DATABASESTORAGE из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetStorageDB().Host, "datahook.cloud.gcm")
			assert.Equal(t, conf.GetStorageDB().Port, 9200)
			assert.Equal(t, conf.GetStorageDB().User, "writer")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD"))
			assert.Equal(t, conf.GetStorageDB().NameDB, "")
			assert.Equal(t, conf.GetStorageDB().Storage.Alert, "test.module_placeholder_alert")
			assert.Equal(t, conf.GetStorageDB().Storage.Case, "test.module_placeholder_case")
		})

		t.Run("Тест 3. Проверка настройки DATABASESTORAGE из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetLogDB().Host, "datahook.cloud.gcm")
			assert.Equal(t, conf.GetLogDB().Port, 9200)
			assert.Equal(t, conf.GetLogDB().User, "log_writer")
			assert.Equal(t, conf.GetLogDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBWLOGPASSWD"))
			assert.Equal(t, conf.GetLogDB().NameDB, "")
			assert.Equal(t, conf.GetLogDB().StorageNameDB, "placeholder_doc-base-db")
		})
	})

	t.Run("Тест чтения переменных окружения", func(t *testing.T) {
		t.Run("Тест 1. Проверка настройки NATS", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDB_NHOST", "127.0.0.1")
			os.Setenv("GO_PHDOCBASEDB_NPORT", "4242")
			os.Setenv("GO_PHDOCBASEDB_NCACHETTL", "650")
			os.Setenv("GO_PHDOCBASEDB_NSUBSENDERCASE", "obj.case")
			os.Setenv("GO_PHDOCBASEDB_NSUBSENDERALERT", "obj.alert")
			os.Setenv("GO_PHDOCBASEDB_NSUBLISTENERCOMMAND", "obj.command")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetNATS().Host, "127.0.0.1")
			assert.Equal(t, conf.GetNATS().Port, 4242)
			assert.Equal(t, conf.GetNATS().CacheTTL, 650)
			assert.Equal(t, conf.GetNATS().Subscriptions.ListenerAlert, "obj.alert")
			assert.Equal(t, conf.GetNATS().Subscriptions.ListenerCase, "obj.case")
			assert.Equal(t, conf.GetNATS().Subscriptions.SenderCommand, "obj.command")
		})

		t.Run("Тест 2. Проверка настройки DATABASESTORAGE", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEHOST", "examle.database.cm")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEPORT", "9559")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGENAME", "any_name")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEUSER", "any_user")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD", "my_new_passwd")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGENALERT", "any.base.alert")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGENCASE", "any.base.case")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetStorageDB().Host, "examle.database.cm")
			assert.Equal(t, conf.GetStorageDB().Port, 9559)
			assert.Equal(t, conf.GetStorageDB().NameDB, "any_name")
			assert.Equal(t, conf.GetStorageDB().User, "any_user")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD"))
			assert.Equal(t, conf.GetStorageDB().Storage.Alert, "any.base.alert")
			assert.Equal(t, conf.GetStorageDB().Storage.Case, "any.base.case")
		})

		t.Run("Тест 3. Проверка настройки DATABASESTORAGE", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDB_DBWLOGHOST", "domaniname.database.cm")
			os.Setenv("GO_PHDOCBASEDB_DBWLOGPORT", "8989")
			os.Setenv("GO_PHDOCBASEDB_DBWLOGUSER", "somebody_user")
			os.Setenv("GO_PHDOCBASEDB_DBWLOGNAME", "any_name_db")
			os.Setenv("GO_PHDOCBASEDB_DBWLOGPASSWD", "your_passwd")
			os.Setenv("GO_PHDOCBASEDB_DBWLOGSTORAGENAME", "log_storage")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetLogDB().Host, "domaniname.database.cm")
			assert.Equal(t, conf.GetLogDB().Port, 8989)
			assert.Equal(t, conf.GetLogDB().User, "somebody_user")
			assert.Equal(t, conf.GetLogDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBWLOGPASSWD"))
			assert.Equal(t, conf.GetLogDB().NameDB, "any_name_db")
			assert.Equal(t, conf.GetLogDB().StorageNameDB, "log_storage")
		})
	})
}
