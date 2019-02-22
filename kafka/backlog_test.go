package kafka

import (
	"fmt"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewBacklogReporter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	retrieverMock := NewMockBacklogRetriever(ctrl)
	logger := func(name string, count int, err error, shutdown bool) {}
	shutdown := false

	br, err := NewBacklogReporter("testBacklogReporterName", retrieverMock, logger, &shutdown, 200)
	if assert.Nil(t, err) {
		assert.NotNil(t, br)
	}

	_, err = NewBacklogReporter("testBacklogReporterName", nil, logger, &shutdown, 200)
	assert.NotNil(t, err)

	_, err = NewBacklogReporter("testBacklogReporterName", retrieverMock, nil, &shutdown, 200)
	assert.NotNil(t, err)

	_, err = NewBacklogReporter("testBacklogReporterName", retrieverMock, logger, nil, 200)
	assert.NotNil(t, err)
}

func TestBackLogReporter_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testBacklogCount = 1
	var testBacklogReporterName = "testBacklogReporterName"

	retrieverMock := NewMockBacklogRetriever(ctrl)
	retrieverMock.EXPECT().
		GetBacklog().
		Return(testBacklogCount, nil).
		AnyTimes()

	shutdown := false
	var callNum int
	var reporterName string
	var backlogCount int
	var backlogErr error
	var shutdownFlag bool
	logger := func(name string, count int, err error, shutdown bool) {
		reporterName = name
		callNum++
		backlogCount = count
		backlogErr = err
		shutdownFlag = shutdown
	}

	reporter, err := NewBacklogReporter(testBacklogReporterName, retrieverMock, logger, &shutdown, 100)
	if assert.Nil(t, err) {
		go reporter.Run()
	}
	time.Sleep(time.Millisecond * 300)
	if assert.Nil(t, backlogErr) {
		assert.Equal(t, backlogCount, testBacklogCount)
		assert.Equal(t, testBacklogReporterName, reporterName)
		assert.True(t, callNum > 1 && callNum < 4, fmt.Sprintf("callNum is %d", callNum))
		assert.False(t, shutdownFlag)
	}

	shutdown = true
	time.Sleep(time.Millisecond * 200)
	shutdownCallNum := callNum
	if assert.Nil(t, backlogErr) {
		assert.Equal(t, backlogCount, testBacklogCount)
		assert.Equal(t, testBacklogReporterName, reporterName)
		assert.True(t, callNum > 2 && callNum < 6, fmt.Sprintf("callNum is %d", callNum))
		//check last call reported shutdown
		assert.True(t, shutdownFlag)
	}
	time.Sleep(time.Millisecond * 200)
	if assert.Nil(t, backlogErr) {
		//check no more calls after shutdown
		assert.Equal(t, shutdownCallNum, callNum)
	}
}
