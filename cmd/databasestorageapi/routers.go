package databasestorageapi

import (
	"context"
	"errors"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func (dbs *DatabaseStorage) router(ctx context.Context) {
	handlersList := map[string]map[string]func(context.Context, any){
		"handling alert": {
			"add alert": dbs.addAlert,
		},
		"handling case": {
			"add case": dbs.addCase,
		},
		"information handling": {
			"add geoip information":  dbs.addGeoIPInformation,
			"add sensor information": dbs.addSensorInformation,
		},
	}

	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-dbs.chInput:
			strErr := "the handler for the accepted request was not found"

			if _, ok := handlersList[msg.Section]; !ok {
				dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(strErr)).Error())

				continue
			}

			if _, ok := handlersList[msg.Section][msg.Command]; !ok {
				dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(strErr)).Error())

				continue
			}

			go handlersList[msg.Section][msg.Command](ctx, msg.Data)
		}
	}
}
