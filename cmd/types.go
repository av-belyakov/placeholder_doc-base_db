package main

import (
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/natsapi"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

type ApplicationRouter struct {
	counter       *countermessage.CounterMessage
	logger        interfaces.Logger
	chToNatsApi   chan<- natsapi.SettingsChanInput
	chFromNatsApi <-chan natsapi.SettingsChanOutput
	chToDBSApi    chan<- databasestorageapi.SettingsChanInput
	chFromDBSApi  <-chan databasestorageapi.SettingsChanOutput
}

type ApplicationRouterSettings struct {
	ChanToNats   chan<- natsapi.SettingsChanInput
	ChanFromNats <-chan natsapi.SettingsChanOutput
	ChanToDBS    chan<- databasestorageapi.SettingsChanInput
	ChanFromDBS  <-chan databasestorageapi.SettingsChanOutput
}
