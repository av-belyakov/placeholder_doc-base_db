package interfaces

import "context"

//**************** счётчик *****************

type Counter interface {
	SendMessage(string, int)
}

//************** логирование ***************

type Logger interface {
	GetChan() <-chan Messager
	Send(msgType, msgData string)
}

type Messager interface {
	GetType() string
	SetType(v string)
	GetMessage() string
	SetMessage(v string)
}

type WriterLoggingData interface {
	Write(typeLogFile, str string) bool
}

//************ каналы *************

type CustomJsonDecoder interface {
	GetterCustomJsonDecode
	SetterCustomJsonDecode
}

type GetterCustomJsonDecode interface {
	GetValue() any
	GetUUID() string
	GetFieldName() string
	GetValueType() string
	GetFieldBranch() string
}

type SetterCustomJsonDecode interface {
	SetValue(any)
	SetUUID(string)
	SetFieldName(string)
	SetValueType(string)
	SetFieldBranch(string)
}

type ChannelResponser interface {
	RequestIdHandler
	GetStatusCode() int
	SetStatusCode(int)
	GetError() error
	SetError(error)
	GetData() []byte
	SetData([]byte)
}

type ChannelRequester interface {
	RequestIdHandler
	CommandHandler
	ElementTypeHandler
	RootIdHandler
	CaseIdHandler
	OrderHandler
	GetData() interface{}
	SetData(interface{})
	GetContext() context.Context
	SetContext(v context.Context)
	GetChanOutput() chan ChannelResponser
	SetChanOutput(chan ChannelResponser)
}

type CaseIdHandler interface {
	GetCaseId() string
	SetCaseId(string)
}

type RequestIdHandler interface {
	GetRequestId() string
	SetRequestId(string)
}

type RootIdHandler interface {
	GetRootId() string
	SetRootId(string)
}

type OrderHandler interface {
	GetOrder() string
	SetOrder(string)
}

type ElementTypeHandler interface {
	GetElementType() string
	SetElementType(string)
}

type CommandHandler interface {
	GetCommand() string
	SetCommand(string)
}
