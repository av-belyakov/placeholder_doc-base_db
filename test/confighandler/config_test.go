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
	os.Unsetenv("GO_PHDOCBASEDB_NCOMMAND")
	os.Unsetenv("GO_PHDOCBASEDB_NSUBLISTENER")

	// Настройки доступа к БД в которую будут записыватся alert и case
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEHOST")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPORT")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGENAME")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEUSER")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")
	os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEN")

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
			assert.Equal(t, conf.GetNATS().Command, "object.commandstype.test")

			list := map[string]string{
				"alert": "object.alerttype.test",
				"case":  "object.casetype.test",
			}
			for k, v := range list {
				s, ok := conf.GetNATS().Subscriptions[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 2. Проверка настройки DATABASESTORAGE из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetStorageDB().Host, "datahook.cloud.gcm")
			assert.Equal(t, conf.GetStorageDB().Port, 9200)
			assert.Equal(t, conf.GetStorageDB().User, "writer")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD"))
			assert.Equal(t, conf.GetStorageDB().NameDB, "")

			list := map[string]string{
				"alert": "test.module_placeholderdb_alert",
				"case":  "test.module_placeholderdb_case",
			}
			for k, v := range list {
				s, ok := conf.GetStorageDB().Storage[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 3. Проверка настройки DATABASEWRITELOG из файла config_test.yml", func(t *testing.T) {
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
			os.Setenv("GO_PHDOCBASEDB_NCOMMAND", "obj.command")
			os.Setenv("GO_PHDOCBASEDB_NSUBLISTENER", "alert:object.env.alerttype.test;case:object.env.casetype.test")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetNATS().Host, "127.0.0.1")
			assert.Equal(t, conf.GetNATS().Port, 4242)
			assert.Equal(t, conf.GetNATS().CacheTTL, 650)
			assert.Equal(t, conf.GetNATS().Command, "obj.command")

			list := map[string]string{
				"alert": "object.env.alerttype.test",
				"case":  "object.env.casetype.test",
			}
			for k, v := range list {
				s, ok := conf.GetNATS().Subscriptions[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 2. Проверка настройки DATABASESTORAGE", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEHOST", "examle.database.cm")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEPORT", "9559")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGENAME", "any_name")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEUSER", "any_user")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD", "my_new_passwd")
			os.Setenv("GO_PHDOCBASEDB_DBSTORAGEN", "alert:test.env.module_placeholderdb_alert;case:test.env.module_placeholderdb_case")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetStorageDB().Host, "examle.database.cm")
			assert.Equal(t, conf.GetStorageDB().Port, 9559)
			assert.Equal(t, conf.GetStorageDB().NameDB, "any_name")
			assert.Equal(t, conf.GetStorageDB().User, "any_user")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD"))

			list := map[string]string{
				"alert": "test.env.module_placeholderdb_alert",
				"case":  "test.env.module_placeholderdb_case",
			}
			for k, v := range list {
				s, ok := conf.GetStorageDB().Storage[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 3. Проверка настройки DATABASEWRITELOG", func(t *testing.T) {
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
