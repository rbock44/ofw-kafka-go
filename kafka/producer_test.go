package kafka

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_EncodeMessage(t *testing.T) {
	/*
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		keySchema := NewMockMessageSchema(ctrl)
		keySchema.EXPECT().
			GetEncoder().
			Return(setupEncoder(ctrl, 0)).
			AnyTimes()

		//FIXME currently not used
		keySchema.EXPECT().
			WriteHeader(gomock.Any()).
			Times(0)

		valueSchema := NewMockMessageSchema(ctrl)
		valueSchema.EXPECT().
			GetEncoder().
			Return(setupEncoder(ctrl, 2)).
			AnyTimes()

		//FIXME currently not used
		valueSchema.EXPECT().
			WriteHeader(gomock.Any()).
			Times(0)
	*/
}

func Test_SendKeyValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageProducer(ctrl)
	m.EXPECT().
		SendKeyValue(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewProducer(gomock.Eq("testTopic"), gomock.Eq("testClientID")).
		Return(m, nil)
	SetFrameworkFactory(f)

	producer, err := NewSingleProducer("testTopic", "testClientID")
	assert.Nil(t, err)

	err = producer.SendKeyValue([]byte{}, []byte{})
	assert.Nil(t, err)

	//Test failure
	m = NewMockMessageProducer(ctrl)
	m.EXPECT().
		SendKeyValue(gomock.Any(), gomock.Any()).
		Return(fmt.Errorf("fail"))
	producer.producer = m
	err = producer.SendKeyValue([]byte{}, []byte{})
	assert.NotNil(t, err)
}
