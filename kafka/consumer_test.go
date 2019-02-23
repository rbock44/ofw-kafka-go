package kafka

import (
	"fmt"
	"io"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_ReadMessage_KeyValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		ReadMessage(gomock.Eq(1000), gomock.Any(), gomock.Any()).
		Do(func(timeoutMs int, keyWriter io.Writer, valueWriter io.Writer) {
			keyWriter.Write([]byte{0x0, 0x0, 0x0, 0x0, byte(testSchemaID), 0x8, 0x74, 0x65, 0x73, 0x74})
			valueWriter.Write([]byte{0x0, 0x0, 0x0, 0x0, byte(testSchemaID), 0x8, 0x74, 0x65, 0x73, 0x74})
		}).
		Return(nil).
		AnyTimes()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID")).
		Return(m, nil)
	SetFrameworkFactory(f)

	consumer, err := NewSingleConsumer("testTopic", "testClientID", setupRegistryMock(t, ctrl, nil, nil, setupDecoder(ctrl)))
	assert.Nil(t, err)

	_, _, err = consumer.ReadMessage(1000)
	assert.Nil(t, err)
}

func Test_ReadMessage_NoMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		ReadMessage(gomock.Eq(1000), gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID")).
		Return(m, nil)
	SetFrameworkFactory(f)

	consumer, err := NewSingleConsumer("testTopic", "testClientID", setupRegistryMock(t, ctrl, nil, nil, setupDecoder(ctrl)))
	assert.Nil(t, err)

	key, value, err := consumer.ReadMessage(1000)
	assert.Equal(t, nil, key)
	assert.Equal(t, nil, value)
	assert.Nil(t, err)
}

func Test_ReadMessage_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		ReadMessage(gomock.Eq(1000), gomock.Any(), gomock.Any()).
		Return(fmt.Errorf("read error")).
		Times(1)

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID")).
		Return(m, nil)
	SetFrameworkFactory(f)

	consumer, err := NewSingleConsumer("testTopic", "testClientID", setupRegistryMock(t, ctrl, nil, nil, setupDecoder(ctrl)))
	assert.Nil(t, err)

	key, value, err := consumer.ReadMessage(1000)
	assert.Equal(t, nil, key)
	assert.Equal(t, nil, value)
	assert.Equal(t, fmt.Errorf("read error"), err)
}

func Test_Process_Shutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var calledShutdown bool
	var calledErr error
	var calledKey interface{}
	var calledValue interface{}
	messageHandler := func(key interface{}, value interface{}, err error) {
		calledKey = key
		calledValue = value
		calledErr = err
	}

	m := NewMockMessageConsumer(ctrl)
	m.EXPECT().
		ReadMessage(gomock.Eq(100), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewConsumer(gomock.Eq("testTopic"), gomock.Eq("testClientID")).
		Return(m, nil)
	SetFrameworkFactory(f)

	shutdown := NewShutdownManager()

	consumer, err := NewBulkConsumer(
		"testTopic",
		"testClientID",
		setupRegistryMock(t, ctrl, nil, nil, setupDecoder(ctrl)),
		messageHandler,
		100, &shutdown.ShutdownState)
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
}

func Test_Process_NoMessageHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	shutdown := false

	_, err := NewBulkConsumer(
		"testTopic",
		"testClientID",
		setupRegistryMock(t, ctrl, nil, nil, nil),
		nil,
		100,
		&shutdown)
	assert.NotNil(t, err)
}
