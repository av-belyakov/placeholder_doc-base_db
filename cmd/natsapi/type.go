package natsapi

import (
	"github.com/nats-io/nats.go"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// apiNatsSettings настройки для API NATS
type apiNatsModule struct {
	natsConn           *nats.Conn
	logger             interfaces.Logger
	host               string
	nameRegionalObject string
	subscriptions      subscription
	chFromModule       chan SettingsOutputChan
	chToModule         chan SettingsInputChan
	cachettl           int
	port               int
}

type subscription struct {
	listenerCase  string
	listenerAlert string
	senderCommand string
}

// NatsApiOptions функциональные опции
type NatsApiOptions func(*apiNatsModule) error

// SettingsOutputChan канал вывода данных из модуля
type SettingsOutputChan struct {
	Data        []byte
	TaskId      string
	SubjectType string
}

// SettingsInputChan канал для приема данных в модуль
type SettingsInputChan struct {
	Command string
	TaskId  string
	CaseId  string
	RootId  string
}
