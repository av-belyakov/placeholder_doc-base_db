package databasestorageapi_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
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

	newGeoIpList := []documentgenerator.IpAddressInformation{
		{Ip: "45.13.191.34", City: "Oslolo", Country: "Норвегия", CountryCode: "NO"},
		{Ip: "45.13.191.28", City: "Oslolo", Country: "Норвегия", CountryCode: "NO"},
		{Ip: "45.13.191.18", City: "Oslo", Country: "Норвегия", CountryCode: "NO"},
		{Ip: "45.13.191.47", City: "Oslo", Country: "Норвегия", CountryCode: "NO"},
		{Ip: "98.113.101.17", City: "Xmu", Country: "Нигерия", CountryCode: "NG"},
	}

	//underlineId, err := apiDBS.SearchUnderlineIdCase(ctx, "module_placeholderdb_case_2025_7", "~88190333152")
	underlineId, geoIpList, err := apiDBS.SearchGeoIPInformationCase(ctx, "module_placeholderdb_case_2025_7", "~1665007800")
	assert.NoError(t, err)
	assert.Equal(t, underlineId, "4g0WIJcBRHX25kGeUOOR")

	t.Log("underlineId:", underlineId)
	fmt.Println("BEFORE @ipAddressAdditionalInformation:")
	for k, v := range geoIpList {
		fmt.Printf("%d. %#v\n", k, v)
	}

	for k, v := range newGeoIpList {
		num, ok := supportingfunctions.SliceContainsElementFunc(geoIpList, func(num int) bool {
			if v.Ip == geoIpList[num].Ip {
				return true
			}

			return false
		})
		if ok {
			geoIpList[num] = v
		} else {
			geoIpList = append(geoIpList, v)
		}

		t.Logf("ok:%t, k:%d, num:%d\n", ok, k, num)
	}

	fmt.Println("AFTER @ipAddressAdditionalInformation:")
	for k, v := range geoIpList {
		fmt.Printf("%d. %#v\n", k, v)
	}

	t.Cleanup(func() {
		os.Unsetenv("GO_PHDOCBASEDB_DBWLOGPASSWD")
		os.Unsetenv("GO_PHDOCBASEDB_DBSTORAGEPASSWD")

		cancel()
	})
}
