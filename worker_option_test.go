package gue

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/vgarvardt/gue/v3/adapter"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Debug(msg string, fields ...adapter.Field) {
	m.Called(msg, fields)
}

func (m *mockLogger) Info(msg string, fields ...adapter.Field) {
	m.Called(msg, fields)
}

func (m *mockLogger) Error(msg string, fields ...adapter.Field) {
	m.Called(msg, fields)
}

func (m *mockLogger) With(fields ...adapter.Field) adapter.Logger {
	args := m.Called(fields)
	return args.Get(0).(adapter.Logger)
}

func TestWithWorkerPollInterval(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerWithDefaultInterval := NewWorker(nil, wm)
	assert.Equal(t, defaultPollInterval, workerWithDefaultInterval.interval)

	customInterval := 12345 * time.Millisecond
	workerWithCustomInterval := NewWorker(nil, wm, WithWorkerPollInterval(customInterval))
	assert.Equal(t, customInterval, workerWithCustomInterval.interval)
}

func TestWithWorkerQueue(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerWithDefaultQueue := NewWorker(nil, wm)
	assert.Equal(t, defaultQueueName, workerWithDefaultQueue.queue)

	customQueue := "fooBarBaz"
	workerWithCustomQueue := NewWorker(nil, wm, WithWorkerQueue(customQueue))
	assert.Equal(t, customQueue, workerWithCustomQueue.queue)
}

func TestWithWorkerID(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerWithDefaultID := NewWorker(nil, wm)
	assert.NotEmpty(t, workerWithDefaultID.id)

	customID := "some-meaningful-id"
	workerWithCustomID := NewWorker(nil, wm, WithWorkerID(customID))
	assert.Equal(t, customID, workerWithCustomID.id)
}

func TestWithWorkerLogger(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerWithDefaultLogger := NewWorker(nil, wm)
	assert.IsType(t, adapter.NoOpLogger{}, workerWithDefaultLogger.logger)

	logMessage := "hello"

	l := new(mockLogger)
	l.On("Info", logMessage, mock.Anything)
	// worker sets id as default logger field
	l.On("With", mock.Anything).Return(l)

	workerWithCustomLogger := NewWorker(nil, wm, WithWorkerLogger(l))
	workerWithCustomLogger.logger.Info(logMessage)

	l.AssertExpectations(t)
}

func TestWithWorkerNextScheduledPollStrategy(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}
	workerWithNextScheduledPollStrategy := NewWorker(nil, wm, WithWorkerNextScheduledPollStrategy())
	assert.Equal(t, nextScheduledPollStrategy, workerWithNextScheduledPollStrategy.pollStrategy)
}

func TestSetWorkerPollStrategy(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}
	workerWithNextScheduledPollStrategy := NewWorker(nil, wm, setWorkerPollStrategy(nextScheduledPollStrategy))
	assert.Equal(t, nextScheduledPollStrategy, workerWithNextScheduledPollStrategy.pollStrategy)

	workerWithDefaultPollStrategy := NewWorker(nil, wm, setWorkerPollStrategy(defaultPollStrategy))
	assert.Equal(t, defaultPollStrategy, workerWithDefaultPollStrategy.pollStrategy)
}

func TestWithPoolPollInterval(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerPoolWithDefaultInterval := NewWorkerPool(nil, wm, 2)
	assert.Equal(t, defaultPollInterval, workerPoolWithDefaultInterval.interval)

	customInterval := 12345 * time.Millisecond
	workerPoolWithCustomInterval := NewWorkerPool(nil, wm, 2, WithPoolPollInterval(customInterval))
	assert.Equal(t, customInterval, workerPoolWithCustomInterval.interval)
}

func TestWithPoolQueue(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerPoolWithDefaultQueue := NewWorkerPool(nil, wm, 2)
	assert.Equal(t, defaultQueueName, workerPoolWithDefaultQueue.queue)

	customQueue := "fooBarBaz"
	workerPoolWithCustomQueue := NewWorkerPool(nil, wm, 2, WithPoolQueue(customQueue))
	assert.Equal(t, customQueue, workerPoolWithCustomQueue.queue)
}

func TestWithPoolID(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerPoolWithDefaultID := NewWorkerPool(nil, wm, 2)
	assert.NotEmpty(t, workerPoolWithDefaultID.id)

	customID := "some-meaningful-id"
	workerPoolWithCustomID := NewWorkerPool(nil, wm, 2, WithPoolID(customID))
	assert.Equal(t, customID, workerPoolWithCustomID.id)
}

func TestWithPoolLogger(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}

	workerPoolWithDefaultLogger := NewWorkerPool(nil, wm, 2)
	assert.IsType(t, adapter.NoOpLogger{}, workerPoolWithDefaultLogger.logger)

	logMessage := "hello"

	l := new(mockLogger)
	l.On("Info", logMessage, mock.Anything)
	// worker pool sets id as default logger field
	l.On("With", mock.Anything).Return(l)

	workerPoolWithCustomLogger := NewWorkerPool(nil, wm, 2, WithPoolLogger(l))
	workerPoolWithCustomLogger.logger.Info(logMessage)

	l.AssertExpectations(t)
}

func TestWithPoolNextScheduledPollStrategy(t *testing.T) {
	wm := WorkMap{
		"MyJob": func(ctx context.Context, j *Job) error {
			return nil
		},
	}
	workerPoolWithNextScheduledPollStrategy := NewWorkerPool(nil, wm, 2, WithPoolNextScheduledPollStrategy())
	assert.Equal(t, nextScheduledPollStrategy, workerPoolWithNextScheduledPollStrategy.pollStrategy)
}
