package databasestorageapi

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/elastic/go-elasticsearch/v8"
)

type DatabaseStorage struct {
	client   *elasticsearch.Client
	logging  interfaces.Logger
	settings settingsDatabaseStorage
}

type settingsDatabaseStorage struct {
	namedb       string
	storageAlert string
	storageCase  string
	user         string
	passwd       string
	host         string
	port         int
}

type DatabaseStorageOptions func(*DatabaseStorage) error
