package queue

import (
	"context"
	"errors"
)

var _ Worker = (*taskWorker)(nil)

// just for unit testing, don't use it.
type taskWorker struct {
	messages chan QueuedMessage
}

func (w *taskWorker) Run(task QueuedMessage) error {
	if v, ok := task.(Job); ok {
		if v.Task != nil {
			_ = v.Task(context.Background())
		}
	}
	return nil
}

func (w *taskWorker) Shutdown() error {
	close(w.messages)
	return nil
}

func (w *taskWorker) Queue(job QueuedMessage) error {
	select {
	case w.messages <- job:
		return nil
	default:
		return errors.New("max capacity reached")
	}
}

func (w *taskWorker) Request() (QueuedMessage, error) {
	select {
	case task := <-w.messages:
		return task, nil
	default:
		return nil, ErrNoTaskInQueue
	}
}
