package kafka

import (
	"bytes"
	"time"
)

//SimpleProducer combines a kafka producer with avro schema support
type SimpleProducer struct {
	Registry    Registry
	Producer    MessageProducer
	RateLimiter *RateLimiter
}

//NewSimpleProducer creates a SimpleProducer
func NewSimpleProducer(messageProducer MessageProducer, registry Registry) (*SimpleProducer, error) {
	simpleProducer := SimpleProducer{}
	simpleProducer.Registry = registry
	simpleProducer.Producer = messageProducer

	return &simpleProducer, nil
}

//SetRateLimit limits the message send to limit per second
func (p *SimpleProducer) SetRateLimit(limitPerSecond int) {
	p.RateLimiter = &RateLimiter{
		StartTime:      time.Now(),
		LimitPerSecond: int64(limitPerSecond),
	}
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
