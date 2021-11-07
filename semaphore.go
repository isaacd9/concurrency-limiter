package concurrency

import (
	"fmt"

	sync_semaphore "golang.org/x/sync/semaphore"
)

type SyncSemaphore struct {
	*sync_semaphore.Weighted
}

func NewSyncSemaphore(limit int64) *SyncSemaphore {
	return &SyncSemaphore{
		Weighted: sync_semaphore.NewWeighted(limit),
	}
}

var (
	_ Semaphore = &SyncSemaphore{}
)

func (s *SyncSemaphore) TryAcquire() error {
	if s.Weighted.TryAcquire(1) {
		return nil
	}
	return fmt.Errorf("too many requests")
}

func (s *SyncSemaphore) Release() {
	s.Weighted.Release(1)
}
