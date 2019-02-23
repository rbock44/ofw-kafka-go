package kafka

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testLimit = 3
var testAboveLimit = int64(4)

func TestRateLimiter_CheckCounterStartTimeBelowLimit(t *testing.T) {
	now := time.Now()
	rl := NewRateLimiter(testLimit)
	rl.StartTime = now
	idleTime := rl.Check(now, 0)

	assert.Equal(t, time.Duration(0), idleTime)
}

func TestRateLimiter_CheckCounterStartTimeAboveLimit(t *testing.T) {
	now := time.Now()
	rl := NewRateLimiter(testLimit)
	rl.StartTime = now
	idleTime := rl.Check(now, testAboveLimit)
	assert.Equal(t, time.Second, idleTime)
}

func TestRateLimiter_CheckCounterRemainingTime20msAboveLimit(t *testing.T) {
	now := time.Now()
	rl := NewRateLimiter(testLimit)
	rl.StartTime = now
	idleTime := rl.Check(now.Add(time.Second-time.Millisecond*20), testAboveLimit)
	assert.Equal(t, time.Millisecond*20, idleTime)
}

func TestRateLimiter_CheckCounterAboveTime(t *testing.T) {
	now := time.Now()
	rl := NewRateLimiter(testLimit)
	rl.StartTime = now
	idleTime := rl.Check(now.Add(time.Second+time.Millisecond*20), testAboveLimit)
	assert.Equal(t, time.Duration(0), idleTime)
}

type testRateCounter struct{ Counter int64 }

func (c *testRateCounter) GetRateCounter() *int64 { return &c.Counter }

func TestRateReporter_calculate(t *testing.T) {
	logger := func(name string, rate float64) {}
	shutdown := false
	rateCounter := int64(0)
	rp, err := NewRateReporter("testRateReporter", &rateCounter, &shutdown, logger, 200)
	if assert.Nil(t, err) {
		assert.Equal(t, float64(0.0), rp.calculateRatePerSecond(0, 0))
		//200ms 100 messages is 500 messages in a second
		assert.Equal(t, float64(500.0), rp.calculateRatePerSecond(200, 100))
	}

	rp, err = NewRateReporter("testRateReporter", &rateCounter, &shutdown, logger, 2000)
	if assert.Nil(t, err) {
		assert.Equal(t, float64(0.0), rp.calculateRatePerSecond(0, 0))
		//2 seconds 100 messages is 50 messages in a second
		assert.Equal(t, float64(50.0), rp.calculateRatePerSecond(200, 100))
	}
}
func TestRateReporter_NewRateReporter(t *testing.T) {
	logger := func(name string, rate float64) {}
	shutdown := false
	rateCounter := int64(100)
	_, err := NewRateReporter("testRateReporter", nil, &shutdown, logger, 200)
	assert.NotNil(t, err)
	_, err = NewRateReporter("testRateReporter", &rateCounter, nil, logger, 200)
	assert.NotNil(t, err)
	_, err = NewRateReporter("testRateReporter", &rateCounter, &shutdown, nil, 200)
	assert.NotNil(t, err)
	rr, err := NewRateReporter("testRateReporter", &rateCounter, &shutdown, logger, 200)
	if assert.Nil(t, err) {
		assert.NotNil(t, rr)
	}
}

func TestRateReporter_Run(t *testing.T) {
	shutdown := false
	var reportedShutdown bool
	var reportedName string
	rateCounter := int64(0)
	rp, err := NewRateReporter("testRateReporter", &rateCounter, &shutdown, func(name string, rate float64) {
		reportedName = name
		reportedShutdown = shutdown
	}, 200)
	if assert.Nil(t, err) {
		go rp.Run()
		rateCounter = int64(200)
		time.Sleep(time.Millisecond * 300)
		assert.Equal(t, "testRateReporter", reportedName)
		assert.False(t, reportedShutdown)
		shutdown = true
		time.Sleep(time.Millisecond * 200)
		assert.True(t, reportedShutdown)
	}
}
