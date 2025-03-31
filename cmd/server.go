package main

import (
	"context"
	"log"
	"os"

	"github.com/av-belyakov/placeholder_doc-base_db/constants"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/appversion"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func server(ctx context.Context) {
	var nameRegionalObject string
	if os.Getenv("GO_PHMISP_MAIN") == "development" {
		nameRegionalObject = "gcm-dev"
	} else if os.Getenv("GO_PHMISP_MAIN") == "test" {
		nameRegionalObject = "gcm-test"
	} else {
		nameRegionalObject = "gcm"
	}

	version, err := appversion.GetAppVersion()
	if err != nil {
		log.Println(err)
	}

	rootPath, err := supportingfunctions.GetRootPath(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
	}

	// ****************************************************************************
	// *********** инициализируем модуль чтения конфигурационного файла ***********
	conf, err := confighandler.New(rootPath)
	if err != nil {
		log.Fatalf("error module 'confighandler': %v", err)
	}
}
