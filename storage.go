package shipwreck

import (
	"fmt"
	"sync"
)

type Storage[T any] interface {
	Append(command ...T) (int64, error)
	Get(start int64, end int64) ([]T, error)
	// Commits the logs to the offset and returns the commited logs
	Commit(offset int64) ([]T, error)
	Discard(start int64, end int64) error

	Length() int64
	Commited() int64
}

var _ Storage[any] = (*MemoryStorage[any])(nil)

type MemoryStorage[T any] struct {
	logsLock sync.Mutex
	logs     []T
	commited []T
}

func (ms *MemoryStorage[T]) Append(command ...T) (int64, error) {
	ms.logsLock.Lock()
	defer ms.logsLock.Unlock()

	ms.logs = append(ms.logs, command...)

	return int64(len(ms.logs)), nil
}

func (ms *MemoryStorage[T]) Commit(offset int64) ([]T, error) {
	ms.logsLock.Lock()
	defer ms.logsLock.Unlock()
	if len(ms.commited) > int(offset) {
		return nil, fmt.Errorf("Invalid state of storage")
	}
	commited := ms.logs[len(ms.commited):offset]
	ms.commited = append(ms.commited, commited...)
	return commited, nil
}

func (ms *MemoryStorage[T]) Get(start int64, end int64) ([]T, error) {
	ms.logsLock.Lock()
	defer ms.logsLock.Unlock()

	return ms.logs[start:end], nil
}

func (ms *MemoryStorage[T]) Length() int64 {
	ms.logsLock.Lock()
	defer ms.logsLock.Unlock()

	return int64(len(ms.logs))
}

func (ms *MemoryStorage[T]) Discard(start int64, end int64) error {
	ms.logsLock.Lock()
	defer ms.logsLock.Unlock()

	ms.logs = append(ms.logs[:start], ms.logs[end:]...)

	return nil
}

func (ms *MemoryStorage[T]) Commited() int64 {
	return int64(len(ms.commited))
}
