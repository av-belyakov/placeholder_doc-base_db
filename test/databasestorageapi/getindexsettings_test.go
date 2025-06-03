package databasestorageapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
)

func TestGetIndexSettings(t *testing.T) {
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

	indexSettings, err := apiDBS.GetIndexSetting(ctx, "module_placeholder_new_case_2025_4")
	assert.NoError(t, err)
	assert.NotEmpty(t, len(indexSettings))

	t.Log("Indexes:")
	for k, v := range indexSettings {
		t.Logf("Index:'%s'\n\tValue:'%+v'\n", k, v)
	}

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}
