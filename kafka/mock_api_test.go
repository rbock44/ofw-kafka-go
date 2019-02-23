// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package kafka is a generated GoMock package.
package kafka

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageSchema is a mock of MessageSchema interface
type MockMessageSchema struct {
	ctrl     *gomock.Controller
	recorder *MockMessageSchemaMockRecorder
}

// MockMessageSchemaMockRecorder is the mock recorder for MockMessageSchema
type MockMessageSchemaMockRecorder struct {
	mock *MockMessageSchema
}

// NewMockMessageSchema creates a new mock instance
func NewMockMessageSchema(ctrl *gomock.Controller) *MockMessageSchema {
	mock := &MockMessageSchema{ctrl: ctrl}
	mock.recorder = &MockMessageSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageSchema) EXPECT() *MockMessageSchemaMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockMessageSchema) GetID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockMessageSchemaMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockMessageSchema)(nil).GetID))
}

// GetDecoder mocks base method
func (m *MockMessageSchema) GetDecoder() Decoder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecoder")
	ret0, _ := ret[0].(Decoder)
	return ret0
}

// GetDecoder indicates an expected call of GetDecoder
func (mr *MockMessageSchemaMockRecorder) GetDecoder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecoder", reflect.TypeOf((*MockMessageSchema)(nil).GetDecoder))
}

// GetEncoder mocks base method
func (m *MockMessageSchema) GetEncoder() Encoder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEncoder")
	ret0, _ := ret[0].(Encoder)
	return ret0
}

// GetEncoder indicates an expected call of GetEncoder
func (mr *MockMessageSchemaMockRecorder) GetEncoder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncoder", reflect.TypeOf((*MockMessageSchema)(nil).GetEncoder))
}

// WriteHeader mocks base method
func (m *MockMessageSchema) WriteHeader(writer io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", writer)
}

// WriteHeader indicates an expected call of WriteHeader
func (mr *MockMessageSchemaMockRecorder) WriteHeader(writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockMessageSchema)(nil).WriteHeader), writer)
}

// MockDecoder is a mock of Decoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockDecoder) Decode(reader io.Reader) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", reader)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockDecoderMockRecorder) Decode(reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), reader)
}

// MockEncoder is a mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Encode mocks base method
func (m *MockEncoder) Encode(value interface{}, writer io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Encode", value, writer)
}

// Encode indicates an expected call of Encode
func (mr *MockEncoderMockRecorder) Encode(value, writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), value, writer)
}

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockRegistry) Register(name string, version int, schemaFile string, decoder Decoder, encoder Encoder) (MessageSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", name, version, schemaFile, decoder, encoder)
	ret0, _ := ret[0].(MessageSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockRegistryMockRecorder) Register(name, version, schemaFile, decoder, encoder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistry)(nil).Register), name, version, schemaFile, decoder, encoder)
}

// GetSchemaByName mocks base method
func (m *MockRegistry) GetSchemaByName(name string) (MessageSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchemaByName", name)
	ret0, _ := ret[0].(MessageSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchemaByName indicates an expected call of GetSchemaByName
func (mr *MockRegistryMockRecorder) GetSchemaByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchemaByName", reflect.TypeOf((*MockRegistry)(nil).GetSchemaByName), name)
}

// GetSchemaByID mocks base method
func (m *MockRegistry) GetSchemaByID(id int) (MessageSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchemaByID", id)
	ret0, _ := ret[0].(MessageSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchemaByID indicates an expected call of GetSchemaByID
func (mr *MockRegistryMockRecorder) GetSchemaByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchemaByID", reflect.TypeOf((*MockRegistry)(nil).GetSchemaByID), id)
}

// MockMessageConsumer is a mock of MessageConsumer interface
type MockMessageConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageConsumerMockRecorder
}

// MockMessageConsumerMockRecorder is the mock recorder for MockMessageConsumer
type MockMessageConsumerMockRecorder struct {
	mock *MockMessageConsumer
}

// NewMockMessageConsumer creates a new mock instance
func NewMockMessageConsumer(ctrl *gomock.Controller) *MockMessageConsumer {
	mock := &MockMessageConsumer{ctrl: ctrl}
	mock.recorder = &MockMessageConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageConsumer) EXPECT() *MockMessageConsumerMockRecorder {
	return m.recorder
}

// ReadMessage mocks base method
func (m *MockMessageConsumer) ReadMessage(timeoutMs int, keyWriter, valueWriter io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage", timeoutMs, keyWriter, valueWriter)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadMessage indicates an expected call of ReadMessage
func (mr *MockMessageConsumerMockRecorder) ReadMessage(timeoutMs, keyWriter, valueWriter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockMessageConsumer)(nil).ReadMessage), timeoutMs, keyWriter, valueWriter)
}

// GetMessageCounter mocks base method
func (m *MockMessageConsumer) GetMessageCounter() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageCounter")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetMessageCounter indicates an expected call of GetMessageCounter
func (mr *MockMessageConsumerMockRecorder) GetMessageCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageCounter", reflect.TypeOf((*MockMessageConsumer)(nil).GetMessageCounter))
}

// GetBacklog mocks base method
func (m *MockMessageConsumer) GetBacklog() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBacklog")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBacklog indicates an expected call of GetBacklog
func (mr *MockMessageConsumerMockRecorder) GetBacklog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBacklog", reflect.TypeOf((*MockMessageConsumer)(nil).GetBacklog))
}

// Close mocks base method
func (m *MockMessageConsumer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMessageConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageConsumer)(nil).Close))
}

// MockMessageProducer is a mock of MessageProducer interface
type MockMessageProducer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageProducerMockRecorder
}

// MockMessageProducerMockRecorder is the mock recorder for MockMessageProducer
type MockMessageProducerMockRecorder struct {
	mock *MockMessageProducer
}

// NewMockMessageProducer creates a new mock instance
func NewMockMessageProducer(ctrl *gomock.Controller) *MockMessageProducer {
	mock := &MockMessageProducer{ctrl: ctrl}
	mock.recorder = &MockMessageProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageProducer) EXPECT() *MockMessageProducerMockRecorder {
	return m.recorder
}

// SendKeyValue mocks base method
func (m *MockMessageProducer) SendKeyValue(key, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendKeyValue", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendKeyValue indicates an expected call of SendKeyValue
func (mr *MockMessageProducerMockRecorder) SendKeyValue(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendKeyValue", reflect.TypeOf((*MockMessageProducer)(nil).SendKeyValue), key, value)
}

// Close mocks base method
func (m *MockMessageProducer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMessageProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageProducer)(nil).Close))
}

// MockBacklogRetriever is a mock of BacklogRetriever interface
type MockBacklogRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockBacklogRetrieverMockRecorder
}

// MockBacklogRetrieverMockRecorder is the mock recorder for MockBacklogRetriever
type MockBacklogRetrieverMockRecorder struct {
	mock *MockBacklogRetriever
}

// NewMockBacklogRetriever creates a new mock instance
func NewMockBacklogRetriever(ctrl *gomock.Controller) *MockBacklogRetriever {
	mock := &MockBacklogRetriever{ctrl: ctrl}
	mock.recorder = &MockBacklogRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBacklogRetriever) EXPECT() *MockBacklogRetrieverMockRecorder {
	return m.recorder
}

// GetBacklog mocks base method
func (m *MockBacklogRetriever) GetBacklog() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBacklog")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBacklog indicates an expected call of GetBacklog
func (mr *MockBacklogRetrieverMockRecorder) GetBacklog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBacklog", reflect.TypeOf((*MockBacklogRetriever)(nil).GetBacklog))
}

// MockDeliveredCounter is a mock of DeliveredCounter interface
type MockDeliveredCounter struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveredCounterMockRecorder
}

// MockDeliveredCounterMockRecorder is the mock recorder for MockDeliveredCounter
type MockDeliveredCounterMockRecorder struct {
	mock *MockDeliveredCounter
}

// NewMockDeliveredCounter creates a new mock instance
func NewMockDeliveredCounter(ctrl *gomock.Controller) *MockDeliveredCounter {
	mock := &MockDeliveredCounter{ctrl: ctrl}
	mock.recorder = &MockDeliveredCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeliveredCounter) EXPECT() *MockDeliveredCounterMockRecorder {
	return m.recorder
}

// GetDeliveredCounter mocks base method
func (m *MockDeliveredCounter) GetDeliveredCounter() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveredCounter")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeliveredCounter indicates an expected call of GetDeliveredCounter
func (mr *MockDeliveredCounterMockRecorder) GetDeliveredCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveredCounter", reflect.TypeOf((*MockDeliveredCounter)(nil).GetDeliveredCounter))
}

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// NewConsumer mocks base method
func (m *MockProvider) NewConsumer(topic, clientID string) (MessageConsumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConsumer", topic, clientID)
	ret0, _ := ret[0].(MessageConsumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConsumer indicates an expected call of NewConsumer
func (mr *MockProviderMockRecorder) NewConsumer(topic, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsumer", reflect.TypeOf((*MockProvider)(nil).NewConsumer), topic, clientID)
}

// NewProducer mocks base method
func (m *MockProvider) NewProducer(topic, clientID string) (MessageProducer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProducer", topic, clientID)
	ret0, _ := ret[0].(MessageProducer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProducer indicates an expected call of NewProducer
func (mr *MockProviderMockRecorder) NewProducer(topic, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProducer", reflect.TypeOf((*MockProvider)(nil).NewProducer), topic, clientID)
}

// NewRegistry mocks base method
func (m *MockProvider) NewRegistry() (Registry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegistry")
	ret0, _ := ret[0].(Registry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRegistry indicates an expected call of NewRegistry
func (mr *MockProviderMockRecorder) NewRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegistry", reflect.TypeOf((*MockProvider)(nil).NewRegistry))
}
