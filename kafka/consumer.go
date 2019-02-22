package kafka

import (
	"bytes"
	"fmt"
)

//SimpleConsumer only supports ReadMessage
type SimpleConsumer struct {
	Consumer MessageConsumer
	Registry Registry
}

//BulkConsumer supports bulk reads with message handlere and error handler
type BulkConsumer struct {
	SimpleConsumer *SimpleConsumer
	MessageHandler func(key interface{}, value interface{}, err error)
	Registry       Registry
	PollTimeMs     int
	Shutdown       *bool
}

//NewSimpleConsumer creates a SimpleConsumer
func NewSimpleConsumer(messageConsumer MessageConsumer, registry Registry) (*SimpleConsumer, error) {
	sc := SimpleConsumer{
		Consumer: messageConsumer,
		Registry: registry,
	}

	return &sc, nil
}

//NewBulkConsumer creates a new BulkConsumer
func NewBulkConsumer(
	messageHandler func(key interface{}, value interface{}, err error),
	simpleConsumer *SimpleConsumer,
	pollTimeMs int,
	shutdown *bool) (*BulkConsumer, error) {
	if messageHandler == nil {
		return nil, fmt.Errorf("messageHandler is nil")
	}
	if simpleConsumer == nil {
		return nil, fmt.Errorf("simpleConsumer is nil")
	}
	bulkConsumer := BulkConsumer{
		MessageHandler: messageHandler,
		SimpleConsumer: simpleConsumer,
		PollTimeMs:     pollTimeMs,
		Shutdown:       shutdown,
	}

	return &bulkConsumer, nil
}

//ReadMessage read a message and process it
func (c *SimpleConsumer) ReadMessage(shutdownCheckInterfaceMs int) (interface{}, interface{}, error) {
	keyBuffer := &bytes.Buffer{}
	valueBuffer := &bytes.Buffer{}
	err := c.Consumer.ReadMessage(shutdownCheckInterfaceMs, keyBuffer, valueBuffer)
	if err != nil {
		return nil, nil, err
	}

	if keyBuffer.Len() == 0 {
		//no message poll interval expired
		return nil, nil, nil
	}

	schemaID, err := readSchemaID(keyBuffer)
	if err != nil {
		return nil, nil, err
	}
	keySchema, err := c.Registry.GetSchemaByID(int(schemaID))
	if err != nil {
		return nil, nil, err
	}

	decoder := keySchema.GetDecoder()
	if decoder == nil {
		return nil, nil, fmt.Errorf("no key decoder")
	}

	key, err := decoder.Decode(keyBuffer)
	if err != nil {
		return nil, nil, err
	}

	id, err := readSchemaID(valueBuffer)
	if err != nil {
		return nil, nil, err
	}
	valueSchema, err := c.Registry.GetSchemaByID(int(id))
	if err != nil {
		return nil, nil, err
	}

	decoder = valueSchema.GetDecoder()
	if decoder == nil {
		return key, nil, fmt.Errorf("no value decoder")
	}
	value, err := decoder.Decode(valueBuffer)

	return key, value, err
}

//Process receive messages and dispatch to message handler and error handler until shutdown flag is true
func (c *BulkConsumer) Process() {
	for {
		key, value, err := c.SimpleConsumer.ReadMessage(c.PollTimeMs)
		if err != nil {
			c.MessageHandler(key, value, err)
		} else {
			if key != nil {
				c.MessageHandler(key, value, err)
			}
		}

		if *c.Shutdown {
			return
		}
	}
}
