package kafka

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

//AvroSchema handles schema lookup and codec
type AvroSchema struct {
	ID      uint32
	Version int
	Subject string
	Content string
	Decoder Decoder
	Encoder Encoder
}

//GetID get the schema ID
func (a AvroSchema) GetID() int {
	return int(a.ID)
}

//GetDecoder get the schema decoder
func (a AvroSchema) GetDecoder() Decoder {
	return a.Decoder
}

//GetEncoder get the schema encoder
func (a AvroSchema) GetEncoder() Encoder {
	return a.Encoder
}

//WriteHeader writes the encoding schema id with magic byte
func (a AvroSchema) WriteHeader(writer io.Writer) {
	header := make([]byte, 5)
	binary.BigEndian.PutUint32(header[1:5], uint32(a.ID))
	writer.Write(header)
}

//SchemaResolver looks up the schema in a remote registry
type SchemaResolver interface {
	GetSchemaBySubject(subject string, version int) (schemaID int, err error)
	RegisterNewSchema(subject string, content string) (id int, err error)
}

//SchemaRegistry contains all registered schema
type SchemaRegistry struct {
	SchemasByID   map[int]*AvroSchema
	SchemasByName map[string]*AvroSchema
	Resolver      SchemaResolver
}

//NewSchemaRegistry creates a kafka schema registry
func NewSchemaRegistry() (*SchemaRegistry, error) {
	schemaResolver, err := fwFactory.NewSchemaResolver()
	if err != nil {
		return nil, err
	}
	return &SchemaRegistry{
		Resolver:      schemaResolver,
		SchemasByID:   map[int]*AvroSchema{},
		SchemasByName: map[string]*AvroSchema{},
	}, nil
}

//Lookup lookup schema in kafka cluster
func (s SchemaRegistry) Lookup(subject string, version int) (*AvroSchema, error) {
	schemaID, err := s.Resolver.GetSchemaBySubject(subject, version)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("schema registry lookup error [%#v]\n", err))
	}
	localSchema := &AvroSchema{
		ID:      uint32(schemaID),
		Subject: subject,
		Version: version,
	}
	s.SchemasByID[schemaID] = localSchema
	s.SchemasByName[subject] = localSchema

	return localSchema, nil
}

//GetSchemaByID gets the schema by id
func (s SchemaRegistry) GetSchemaByID(id int) (MessageSchema, error) {
	ms := s.SchemasByID[id]
	if ms == nil {
		return nil, fmt.Errorf("schema id not found [%d]", id)
	}
	return ms, nil
}

//GetSchemaByName gets the schema by name
func (s SchemaRegistry) GetSchemaByName(name string) (MessageSchema, error) {
	ms := s.SchemasByName[name]
	if ms == nil {
		return nil, fmt.Errorf("schema subject not found [%s]", name)
	}
	return ms, nil
}

//Register schema with registry
func (s SchemaRegistry) Register(subject string, version int, schemaPath string, decoder Decoder, encoder Encoder) (MessageSchema, error) {
	retryCount := 5
	retryTimeMs := 1000
	schemaContent, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read schema file [%s]", schemaPath)
	}
	var id int
	for i := 0; i <= retryCount; i++ {
		id, err = s.Resolver.RegisterNewSchema(subject, string(schemaContent))
		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(retryTimeMs))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot register avro schema [%#v]", err)
	}
	localSchema := &AvroSchema{
		ID:      uint32(id),
		Subject: subject,
		Version: version,
		Decoder: decoder,
		Encoder: encoder,
	}
	s.SchemasByID[id] = localSchema
	s.SchemasByName[subject] = localSchema
	return localSchema, nil
}

//DecodeMessage extracts the schema version and decodes the key and value
func (s SchemaRegistry) DecodeMessage(context *MessageContext, key []byte, value []byte) (interface{}, interface{}, error) {
	keyBuffer := bytes.NewBuffer(key)
	valueBuffer := bytes.NewBuffer(value)

	if keyBuffer.Len() == 0 {
		//no message poll interval expired
		return nil, nil, nil
	}

	schemaID, err := readSchemaID(keyBuffer)
	if err != nil {
		return nil, nil, err
	}
	keySchema, err := s.GetSchemaByID(int(schemaID))
	if err != nil {
		return nil, nil, err
	}

	decoder := keySchema.GetDecoder()
	if decoder == nil {
		return nil, nil, fmt.Errorf("no key decoder")
	}

	decodedKey, err := decoder.Decode(keyBuffer)
	if err != nil {
		return nil, nil, err
	}

	id, err := readSchemaID(valueBuffer)
	if err != nil {
		return nil, nil, err
	}
	valueSchema, err := s.GetSchemaByID(int(id))
	if err != nil {
		return nil, nil, err
	}

	decoder = valueSchema.GetDecoder()
	if decoder == nil {
		return key, nil, fmt.Errorf("no value decoder")
	}
	decodedValue, err := decoder.Decode(valueBuffer)

	return decodedKey, decodedValue, err
}
