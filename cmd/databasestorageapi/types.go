package databasestorageapi

import (
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/elastic/go-elasticsearch/v8"
)

type DatabaseStorage struct {
	counter  interfaces.Counter
	logger   interfaces.Logger
	client   *elasticsearch.Client
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
	Data    any
	Section string
	Command string
}

// SettingsChanOutput параметры канала для приёма данных из модуля
type SettingsChanOutput struct {
	CaseId  string
	RootId  string
	Command string
}

type DatabaseStorageOptions func(*DatabaseStorage) error

// *** для объектов типа 'alert' ***
// AlertDBResponse информация о кейсах
type AlertDBResponse struct {
	Options AlertDBResponseOptions `json:"hits"`
}

// AlertDBResponseOptions информация о кейсах
type AlertDBResponseOptions struct {
	Total    OptionsTotal           `json:"total"`
	MaxScore float64                `json:"max_score"`
	Hits     []PatternVerifiedAlert `json:"hits"`
}

// PatternVerifiedAlert шаблон
type PatternVerifiedAlert struct {
	Source documentgenerator.VerifiedAlert `json:"_source"`
	ServiseOption
}

// *** для объектов типа 'case' ***
// CaseDBResponse информация о кейсах
type CaseDBResponse struct {
	Options CaseDBResponseOptions `json:"hits"`
}

// CaseDBResponseOptions информация о кейсах
type CaseDBResponseOptions struct {
	Total    OptionsTotal          `json:"total"`
	MaxScore float64               `json:"max_score"`
	Hits     []PatternVerifiedCase `json:"hits"`
}

// PatternVerifiedCase шаблон
type PatternVerifiedCase struct {
	Source documentgenerator.VerifiedCase `json:"_source"`
	ServiseOption
}

// OptionsTotal опции в результате поиска
type OptionsTotal struct {
	Relation string `json:"relation"` //отношение (==, >, <)
	Value    int    `json:"value"`    //количество найденных значений
}

// ServiseOption сервисные опции
type ServiseOption struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}

type listSensorId struct {
	sensors []string
}
