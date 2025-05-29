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

	t.Log("underlineId:", underlineId)

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}
