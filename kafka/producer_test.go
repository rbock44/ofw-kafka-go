package kafka

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_SendKeyValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockMessageProducer(ctrl)
	m.EXPECT().
		SendKeyValue(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	keySchema := NewMockMessageSchema(ctrl)
	keySchema.EXPECT().
		GetEncoder().
		Return(setupEncoder(ctrl, 2)).
		AnyTimes()

	keySchema.EXPECT().
		WriteHeader(gomock.Any()).
		Times(2)

	valueSchema := NewMockMessageSchema(ctrl)
	valueSchema.EXPECT().
		GetEncoder().
		Return(setupEncoder(ctrl, 2)).
		AnyTimes()

	valueSchema.EXPECT().
		WriteHeader(gomock.Any()).
		Times(2)

	producer, err := NewSimpleProducer(m, setupRegistryMock(t, ctrl, nil, nil, nil))
	assert.Nil(t, err)

	key := "test"
	err = producer.SendKeyValue(keySchema, key, valueSchema, key)
	assert.Nil(t, err)

	//Test failure
	m = NewMockMessageProducer(ctrl)
	m.EXPECT().
		SendKeyValue(gomock.Any(), gomock.Any()).
		Return(fmt.Errorf("fail"))
	producer.Producer = m
	err = producer.SendKeyValue(keySchema, key, valueSchema, key)
	assert.NotNil(t, err)
}
