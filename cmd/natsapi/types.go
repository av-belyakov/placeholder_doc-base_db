package natsapi

import (
	"github.com/nats-io/nats.go"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// apiNatsSettings настройки для API NATS
type apiNatsModule struct {
	counter       interfaces.Counter
	logger        interfaces.Logger
	natsConn      *nats.Conn
	subscriptions map[string]string
	settings      apiNatsSettings
	chFromModule  chan SettingsChanOutput
	chToModule    chan SettingsChanInput
}

type apiNatsSettings struct {
	nameRegionalObject string
	command            string
	host               string
	cachettl           int
	port               int
}

// NatsApiOptions функциональные опции
type NatsApiOptions func(*apiNatsModule) error

// SettingsChanOutput канал вывода данных из модуля
type SettingsChanOutput struct {
	Data        []byte
	TaskId      string
	SubjectType string
}

// SettingsChanInput канал для приема данных в модуль
type SettingsChanInput struct {
	Data    []byte
	Command string
	TaskId  string
	CaseId  string
	RootId  string
}
