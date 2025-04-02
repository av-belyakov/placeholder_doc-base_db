package confighandler

import "errors"

// GetCommon общие настройки
func (c *ConfigApp) GetCommon() *CfgCommon {
	return &c.Common
}

// GetNATS настройки для NATS
func (c *ConfigApp) GetNATS() *CfgNats {
	return &c.NATS
}

// GetStorageDB настройки для БД отвечающей за хранение данных
func (c *ConfigApp) GetStorageDB() *CfgStorageDB {
	return &c.StorageDB
}

// GetLogDB настройки для БД обеспечивающей логирование данных
func (c *ConfigApp) GetLogDB() *CfgWriteLogDB {
	return &c.LogDB
}

// GetListLogs список типов логирования (error, info и т.д.)
func (c *ConfigApp) GetListLogs() []*LogSet {
	return c.Common.Logs
}

// GetZabbix настройки взаимодействия с Zabbix
func (c *ConfigApp) GetZabbix() *ZabbixOptions {
	return &c.Common.Zabbix
}

// SetNameMessageType наименование типа логирования
func (l *LogSet) SetNameMessageType(v string) error {
	if v == "" {
		return errors.New("the value 'MsgTypeName' must not be empty")
	}

	return nil
}

// SetMaxLogFileSize максимальный размер файла для логирования
func (l *LogSet) SetMaxLogFileSize(v int) error {
	if v < 1000 {
		return errors.New("the value 'MaxFileSize' must not be less than 1000")
	}

	return nil
}

// SetPathDirectory путь к директории логирования
func (l *LogSet) SetPathDirectory(v string) error {
	if v == "" {
		return errors.New("the value 'PathDirectory' must not be empty")
	}

	return nil
}

// SetWritingStdout запись логов на вывод stdout
func (l *LogSet) SetWritingStdout(v bool) {
	l.WritingStdout = v
}

// SetWritingFile запись логов в файл
func (l *LogSet) SetWritingFile(v bool) {
	l.WritingFile = v
}

// SetWritingDB запись логов  в БД
func (l *LogSet) SetWritingDB(v bool) {
	l.WritingDB = v
}

// GetNameMessageType наименование тпа логирования
func (l *LogSet) GetNameMessageType() string {
	return l.MsgTypeName
}

// GetMaxLogFileSize максимальный размер файла для логирования
func (l *LogSet) GetMaxLogFileSize() int {
	return l.MaxFileSize
}

// GetPathDirectory путь к директории логирования
func (l *LogSet) GetPathDirectory() string {
	return l.PathDirectory
}

// GetWritingStdout запись логов на вывод stdout
func (l *LogSet) GetWritingStdout() bool {
	return l.WritingStdout
}

// GetWritingFile запись логов в файл
func (l *LogSet) GetWritingFile() bool {
	return l.WritingFile
}

// GetWritingDB запись логов  в БД
func (l *LogSet) GetWritingDB() bool {
	return l.WritingDB
}
