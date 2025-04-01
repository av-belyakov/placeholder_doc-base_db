package databasestorageapi

import "errors"

//******************* функции настройки опций databasestorageapi ***********************

// WithHost метод устанавливает имя или ip адрес хоста API
func WithHost(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		dbs.settings.host = v

		return nil
	}
}

// WithPort метод устанавливает порт API
func WithPort(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		dbs.settings.port = v

		return nil
	}
}

// WithUser метод устанавливает имя пользователя для доступа к БД
func WithUser(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		dbs.settings.user = v

		return nil
	}
}

// WithPasswd метод устанавливает пароль для доступа к БД
func WithPasswd(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		dbs.settings.passwd = v

		return nil
	}
}

// WithNameDB метод устанавливает наименование БД
func WithNameDB(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.namedb = v

		return nil
	}
}

// WithStorageAlert метод устанавливает наименование коллекции или индекса БД
func WithStorageAlert(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.storageAlert = v

		return nil
	}
}

// WithStorageAlert метод устанавливает наименование коллекции или индекса БД
func WithStorageCase(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.storageAlert = v

		return nil
	}
}
