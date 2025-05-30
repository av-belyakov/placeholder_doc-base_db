package databasestorageapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
)

func TestSearchUnderlineId(t *testing.T) {
	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		t.Log(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	apiDBS, err := ConnectionInitialization(
		ctx,
		databasestorageapi.WithHost("datahook.cloud.gcm"),
		databasestorageapi.WithPort(9200),
		databasestorageapi.WithUser("writer"),
		databasestorageapi.WithPasswd(os.Getenv("GO_PHDOCBASEDB_DBWLOGPASSWD")),
		databasestorageapi.WithStorage(map[string]string{
			"alert": "module_placeholderdb_alert",
			"case":  "module_placeholderdb_case",
		}),
	)
	assert.NoError(t, err)

	underlineId, err := apiDBS.SearchUnderlineIdCase(ctx, "module_placeholder_new_case_2025_5", "~1402712072")
	assert.NoError(t, err)

	/*
			Тест не проходит!!!

		По причине не верного анмаршалинга полученного, из эластика, объекта.
		Навеное надо написать функцию которая с помощью рефлексии разбирала json.
		На вход получала []byte (json) и список имён объектов которые нужно найти
		в этом jsone. На выходе был бы
		map[<имя искомого объекта>]struct{
		 Type string //наименование типа
		 Data any //значение которое можно было бы привести к типу из Type
		}
	*/

	t.Log("underlineId:", underlineId)

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}
