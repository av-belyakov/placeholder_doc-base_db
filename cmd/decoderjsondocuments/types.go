package decoderjsondocuments

import (
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

type HandlerJsonMessageSettings struct {
	logger  interfaces.Logger
	counter *countermessage.CounterMessage
}
