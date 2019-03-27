package kafka

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	Validate func(context *MessageContext, key []byte, value []byte)
}

func (*testHandler) Handle(context *MessageContext, key []byte, value []byte) {

}

func Test_ReadMessage_KeyValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		Process(1000).
		Return(nil).
		/*
			Do(func(timeoutMs int) {
				keyBuffer := &bytes.Buffer{}
				keyBuffer.Write([]byte{0x0, 0x0, 0x0, 0x0, byte(testSchemaID), 0x8, 0x74, 0x65, 0x73, 0x74})
				valueBuffer := &bytes.Buffer{}
				valueBuffer.Write([]byte{0x0, 0x0, 0x0, 0x0, byte(testSchemaID), 0x8, 0x74, 0x65, 0x73, 0x74})
			}).
		*/
		AnyTimes()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID"), &testHandler{}).
		Return(m, nil)
	SetFrameworkFactory(f)

	consumer, err := NewConsumer("testTopic", "testClientID", 1000, &testHandler{})
	assert.Nil(t, err)

	go consumer.Process()
	consumer.Shutdown = true
	assert.Nil(t, err)
}

func Test_Process_Shutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var calledShutdown bool
	var calledErr error
	var calledKey interface{}
	var calledValue interface{}

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		Process(1000).
		Return(nil).
		AnyTimes()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID"), &testHandler{}).
		Return(m, nil)
	SetFrameworkFactory(f)

	shutdown := NewShutdownManager()

	consumer, err := NewConsumer(
		"testTopic",
		"testClientID",
		1000,
		&testHandler{})
	assert.Nil(t, err)

	go consumer.Process()
	time.Sleep(time.Millisecond * 200)
	assert.False(t, calledShutdown)
	assert.Nil(t, calledErr)
	assert.Nil(t, calledKey)
	assert.Nil(t, calledValue)
	shutdown.SignalShutdown()
	time.Sleep(time.Millisecond * 200)
	assert.Nil(t, calledErr)
	assert.Nil(t, calledKey)
	assert.Nil(t, calledValue)
	time.Sleep(1000)
}

func Test_Process_NoMessageHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, err := NewConsumer(
		"testTopic",
		"testClientID",
		1000,
		nil)
	assert.NotNil(t, err)
}
