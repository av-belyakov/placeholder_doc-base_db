package databasestorageapi

import "errors"

// GetChanDataToModule канал для передачи данных в модуль
func (dbs *DatabaseStorage) GetChanDataToModule() chan SettingsChanInput {
	return dbs.chInput
}

// GetChanDataFromModule канал для приёма данных из модуля
func (dbs *DatabaseStorage) GetChanDataFromModule() chan SettingsChanOutput {
	return dbs.chOutput
}

//******************* функции настройки опций databasestorageapi ***********************

// WithHost имя или ip адрес хоста API
func WithHost(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		dbs.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		dbs.settings.port = v

		return nil
	}
}

// WithUser имя пользователя для доступа к БД
func WithUser(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		dbs.settings.user = v

		return nil
	}
}

// WithPasswd пароль для доступа к БД
func WithPasswd(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		dbs.settings.passwd = v

		return nil
	}
}

// WithNameDB наименование БД
func WithNameDB(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.namedb = v

		return nil
	}
}

// WithStorage наименование коллекции или индекса БД
func WithStorage(v map[string]string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.storages = v

		return nil
	}
}

//****************** вспомогательные методы *************************

// Get возвращает список идентификаторов сенсоров
func (e *listSensorId) get() []string {
	return e.sensors
}

// AddElem добавляет только уникальные элементы
func (e *listSensorId) addElem(sensorId string) {
	for _, v := range e.sensors {
		if v == sensorId {
			return
		}
	}

	e.sensors = append(e.sensors, sensorId)
}
