package databasestorageapi

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/elastic/go-elasticsearch/v8"
)

type DatabaseStorage struct {
	client   *elasticsearch.Client
	logging  interfaces.Logger
	settings settingsDatabaseStorage
	chInput  chan SettingsChanInput  //канал для передачи данных в модуль
	chOutput chan SettingsChanOutput //канал для приёма данных из модуля
}

type settingsDatabaseStorage struct {
	storages map[string]string
	namedb   string
	user     string
	passwd   string
	host     string
	port     int
}

// SettingsChanInput параметры канала для передачи данных в модуль
type SettingsChanInput struct {
}

// SettingsChanOutput параметры канала для приёма данных из модуля
type SettingsChanOutput struct {
}

type DatabaseStorageOptions func(*DatabaseStorage) error
