package kafka

import (
	"bytes"
)

//SimpleProducer combines a kafka producer with avro schema support
type SimpleProducer struct {
	Registry Registry
	Producer MessageProducer
}

//NewSimpleProducer creates a SimpleProducer
func NewSimpleProducer(messageProducer MessageProducer, registry Registry) (*SimpleProducer, error) {
	simpleProducer := SimpleProducer{}
	simpleProducer.Registry = registry
	simpleProducer.Producer = messageProducer

	return &simpleProducer, nil
}

//SendKeyValue sends a key value pair
func (p *SimpleProducer) SendKeyValue(keySchema MessageSchema, key interface{}, valueSchema MessageSchema, value interface{}) error {
	keyBuffer := &bytes.Buffer{}
	valueBuffer := &bytes.Buffer{}
	keySchema.WriteHeader(keyBuffer)
	valueSchema.WriteHeader(valueBuffer)
	keySchema.GetEncoder().Encode(key, keyBuffer)
	valueSchema.GetEncoder().Encode(value, valueBuffer)
	err := p.Producer.SendKeyValue(keyBuffer.Bytes(), valueBuffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
