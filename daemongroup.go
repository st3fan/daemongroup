package daemongroup

import (
	"context"
	"sync"
)

type DaemonGroup struct {
	ctx context.Context
	wg  sync.WaitGroup
}

func NewDaemonGroup(ctx context.Context) *DaemonGroup {
	return &DaemonGroup{ctx: ctx}
}

func (dg *DaemonGroup) Go(f func(ctx context.Context)) {
	dg.wg.Add(1)
	go func() {
		f(dg.ctx)
		dg.wg.Done()
	}()
}

func (dg *DaemonGroup) Wait() {
	dg.wg.Wait()
}

