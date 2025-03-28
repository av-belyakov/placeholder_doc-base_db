package confighandler

func (c *ConfigApp) GetCommon() *CfgCommon {
	return &c.Common
}

func (c *ConfigApp) GetNATS() *CfgNats {
	return &c.NATS
}

func (c *ConfigApp) GetStorageDB() *CfgStorageDB {
	return &c.StorageDB
}

func (c *ConfigApp) GetLogDB() *CfgWriteLogDB {
	return &c.LogDB
}
