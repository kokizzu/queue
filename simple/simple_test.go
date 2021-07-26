package simple

import (
	"testing"
	"time"

	"github.com/appleboy/queue"

	"github.com/stretchr/testify/assert"
)

type mockMessage struct {
	msg string
}

func (m mockMessage) Bytes() []byte {
	return []byte(m.msg)
}

func TestQueueUsage(t *testing.T) {
	w := NewWorker()
	assert.Equal(t, defaultQueueSize, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	assert.NoError(t, w.Queue(&mockMessage{}))
	assert.Equal(t, 1, w.Usage())
}

func TestMaxCapacity(t *testing.T) {
	w := NewWorker(WithQueueNum(2))
	assert.Equal(t, 2, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	assert.NoError(t, w.Queue(&mockMessage{}))
	assert.Equal(t, 1, w.Usage())
	assert.NoError(t, w.Queue(&mockMessage{}))
	assert.Equal(t, 2, w.Usage())
	assert.Error(t, w.Queue(&mockMessage{}))
	assert.Equal(t, 2, w.Usage())

	err := w.Queue(&mockMessage{})
	assert.Equal(t, errMaxCapacity, err)
}

func TestCustomFuncAndWait(t *testing.T) {
	m := mockMessage{
		msg: "foo",
	}
	w := NewWorker(
		WithRunFunc(func(msg queue.QueuedMessage, s <-chan struct{}) error {
			time.Sleep(500 * time.Millisecond)
			return nil
		}),
	)
	q, err := queue.NewQueue(
		queue.WithWorker(w),
		queue.WithWorkerCount(2),
	)
	assert.NoError(t, err)
	q.Start()
	time.Sleep(100 * time.Millisecond)
	assert.NoError(t, w.Queue(m))
	assert.NoError(t, w.Queue(m))
	assert.NoError(t, w.Queue(m))
	assert.NoError(t, w.Queue(m))
	time.Sleep(600 * time.Millisecond)
	q.Shutdown()
	q.Wait()
	// you will see the execute time > 1000ms
}
