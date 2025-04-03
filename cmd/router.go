package main

import "context"

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(settings ApplicationRouterSettings) *ApplicationRouter {
	return &ApplicationRouter{
		chToNatsApi:   settings.ChanToNats,
		chFromNatsApi: settings.ChanFromNats,
		chToDBSApi:    settings.ChanToDBS,
		chFromDBSApi:  settings.ChanFromDBS,
	}
}

func (r *ApplicationRouter) Router(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

		case msg := <-r.chFromNatsApi:

		case msg := <-r.chFromDBSApi:

		}
	}
}
