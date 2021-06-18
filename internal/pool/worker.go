package pool

import (
	"context"
)

type worker struct {
	Handler
	ch <-chan string
}

func (w *worker) listen(ctx context.Context) error {
	//resultCh is channel that used locally to send first non nil error from handler
	//or nil when jobs channel is closed and listening ends
	resultCh := make(chan error)

	go func() {
		for url := range w.ch {
			if err := w.Handle(url); err != nil {
				resultCh <- err
				return
			}
		}
		resultCh <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-resultCh:
		return err
	}
}
