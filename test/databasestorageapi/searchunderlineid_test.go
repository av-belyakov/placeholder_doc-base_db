package databasestorageapi_test

import (
	"context"
	"fmt"
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

	fmt.Println("GO_PHDOCBASEDB_DBWLOGPASSWD =", os.Getenv("GO_PHDOCBASEDB_DBWLOGPASSWD"))
	fmt.Println("GO_PHDOCBASEDB_DBSTORAGEPASSWD =", os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD"))

	apiDBS, err := ConnectionInitialization(
		ctx,
		databasestorageapi.WithHost("datahook.cloud.gcm"),
		databasestorageapi.WithPort(9200),
		databasestorageapi.WithUser("writer"),
		databasestorageapi.WithPasswd(os.Getenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")),
		databasestorageapi.WithStorage(map[string]string{
			"alert": "module_placeholderdb_alert",
			"case":  "module_placeholderdb_case",
		}),
	)
	assert.NoError(t, err)

	/*
	   для "module_placeholder_case_2024_7", "~86803587072" нет
	*/

	//underlineId, err := apiDBS.SearchUnderlineIdCase(ctx, "module_placeholderdb_case_2025_7", "~88190333152")
	underlineId, listGeoIp, err := apiDBS.SearchGeoIPInformationCase(ctx, "module_placeholderdb_case_2025_7", "~88190333152")
	assert.NoError(t, err)
	assert.Equal(t, underlineId, "4g0WIJcBRHX25kGeUOOR")

	t.Log("underlineId:", underlineId)
	fmt.Printf("@ipAddressAdditionalInformation:\n%#v\n", listGeoIp)

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}
