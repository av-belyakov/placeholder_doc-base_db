package confighandler

type ConfigApp struct {
	Common    CfgCommon
	NATS      CfgNats
	StorageDB CfgStorageDB
	LogDB     CfgWriteLogDB
}

type CfgCommon struct {
	Logs   []*LogSet
	Zabbix ZabbixOptions
}

type Logs struct {
	Logging []*LogSet
}

type LogSet struct {
	MsgTypeName   string `validate:"oneof=error info warning" yaml:"msgTypeName"`
	PathDirectory string `validate:"required" yaml:"pathDirectory"`
	MaxFileSize   int    `validate:"min=1000" yaml:"maxFileSize"`
	WritingStdout bool   `validate:"required" yaml:"writingStdout"`
	WritingFile   bool   `validate:"required" yaml:"writingFile"`
	WritingDB     bool   `validate:"required" yaml:"writingDB"`
}

type ZabbixSet struct {
	Zabbix ZabbixOptions
}

type ZabbixOptions struct {
	EventTypes  []EventType `yaml:"eventType"`
	NetworkHost string      `validate:"required" yaml:"networkHost"`
	ZabbixHost  string      `validate:"required" yaml:"zabbixHost"`
	NetworkPort int         `validate:"gt=0,lte=65535" yaml:"networkPort"`
}

type EventType struct {
	EventType  string    `validate:"required" yaml:"eventType"`
	ZabbixKey  string    `validate:"required" yaml:"zabbixKey"`
	Handshake  Handshake `yaml:"handshake"`
	IsTransmit bool      `yaml:"isTransmit"`
}

type Handshake struct {
	Message      string `validate:"required" yaml:"message"`
	TimeInterval int    `yaml:"timeInterval"`
}

type CfgNats struct {
	Subscriptions SubscriptionsNATS `yaml:"subscriptions"`
	Host          string            `validate:"required" yaml:"host"`
	Port          int               `validate:"gt=0,lte=65535" yaml:"port"`
	CacheTTL      int               `validate:"gt=10,lte=86400" yaml:"cache_ttl"`
}

type SubscriptionsNATS struct {
	ListenerAlert string `validate:"required" yaml:"listener_alert"`
	ListenerCase  string `validate:"required" yaml:"listener_case"`
	SenderCommand string `validate:"required" yaml:"sender_command"`
}

type CfgStorageDB struct {
	Host    string           `yaml:"host"`
	User    string           `yaml:"user"`
	Passwd  string           `yaml:"passwd"`
	NameDB  string           `yaml:"namedb"`
	Storage CfgStorageNameDB `yaml:"storage_name_db"`
	Port    int              `validate:"gt=0,lte=65535" yaml:"port"`
}

type CfgStorageNameDB struct {
	Alert string `validate:"required" yaml:"alert"`
	Case  string `validate:"required" yaml:"case"`
}

type CfgWriteLogDB struct {
	Host          string `yaml:"host"`
	User          string `yaml:"user"`
	Passwd        string `yaml:"passwd"`
	NameDB        string `yaml:"namedb"`
	StorageNameDB string `yaml:"storage_name_db"`
	Port          int    `validate:"gt=0,lte=65535" yaml:"port"`
}
