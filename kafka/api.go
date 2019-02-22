package kafka

import "io"

//go:generate mockgen -source api.go -destination mock_api_test.go -package kafka MessageConsumer,MessageProducer,MessageSchema,Decoder,Encoder,Registry

//MessageSchema hold the schema of the message
type MessageSchema interface {
	GetID() int
	GetDecoder() Decoder
	GetEncoder() Encoder
	WriteHeader(writer io.Writer)
}

//Decoder decodes the binary stream to a value
type Decoder interface {
	Decode(reader io.Reader) (value interface{}, err error)
}

//Encoder encodes the value to a binary stream
type Encoder interface {
	Encode(value interface{}, writer io.Writer)
}

//Registry abstracts the schema registry
type Registry interface {
	Register(name string, version int, schemaFile string, decoder Decoder, encoder Encoder) (MessageSchema, error)
	GetSchemaByName(name string) (MessageSchema, error)
	GetSchemaByID(id int) (MessageSchema, error)
}

//MessageConsumer interface to abstract message receiving and make writing tests simpler
type MessageConsumer interface {
	ReadMessage(timeoutMs int, keyWriter io.Writer, valueWriter io.Writer) error
}

//MessageProducer interface to abstract message sending and make writing tests simpler
type MessageProducer interface {
	SendKeyValue(key []byte, value []byte) error
}

//BacklogRetriever retrieves the backlog of a consumer
type BacklogRetriever interface {
	GetBacklog() (backlog int, err error)
}
