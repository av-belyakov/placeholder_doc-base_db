package databasestorageapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

const (
	Index_Test         = "test.module_placeholder_case"
	Filepath_Test_Case = "../test_json/case_39100.json"
)

var listIPAddress []databasestorageapi.IpAddressesInformation = []databasestorageapi.IpAddressesInformation{
	{
		Ip:          "96.136.64.9",
		City:        "Havana",
		Country:     "Kuba",
		CountryCode: "CU",
	},
	{
		Ip:          "72.31.66.61",
		City:        "Sidney",
		Country:     "Australia",
		CountryCode: "AU",
	},
	{
		Ip:          "13.22.63.6",
		City:        "Lida",
		Country:     "Livia",
		CountryCode: "LI",
	},
}

func CreateTestCase(ctx context.Context, filePath string) (string, *documentgenerator.VerifiedCase, error) {
	var (
		rootId     string
		verifyCase *documentgenerator.VerifiedCase
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logging := NewLogging()
	counting := countermessage.New(logging.ch)
	counting.Start(ctx)

	b, err := os.ReadFile(filePath)
	if err != nil {
		return rootId, verifyCase, err
	}

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

	rootId, verifyCase, _ = documentgenerator.CaseGenerator(chDecode)

	return rootId, verifyCase, nil
}

func TestUpdateIndexCaseGeoIp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// получаем тестовый кейс
	rootId, verifyCase, err := CreateTestCase(ctx, Filepath_Test_Case)
	if err != nil {
		t.Fatal(err)
	}

	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal(err)
	}

	//подключение к БД
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
	if err != nil {
		t.Fatal(err)
	}

	//удаление тестового индекса
	if err := apiDBS.DelIndexSetting(ctx, []string{Index_Test}); err != nil {
		t.Log("ERROR:", err)
	}

	vcb, err := json.Marshal(verifyCase)
	if err != nil {
		t.Fatal(err)
	}

	//добавляем новый тестовый кейс
	res, err := apiDBS.InsertDocument("", Index_Test, vcb)
	assert.NoError(t, err)

	//проверяем наличие кейса
	underlineId, err := apiDBS.SearchUnderlineIdCase(ctx, Index_Test, rootId)
	assert.NoError(t, err)

	t.Log("underlineId:", underlineId)

	request, err := json.MarshalIndent(databasestorageapi.AdditionalInformationIpAddress{
		IpAddresses: listIPAddress,
	}, "", " ")
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err = apiDBS.Update(Index_Test, underlineId, bodyUpdate)
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()

		res.Body.Close()
	})
}

/*
	t.Run("", func(t *testing.T) {

	})
*/
