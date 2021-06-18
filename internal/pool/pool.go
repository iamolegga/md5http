package pool

import (
	"context"
	"sync"
)

type Pool struct {
	//ch is used to pass job to workers
	ch chan string
	//errCh is used to pass first non nil error from any worker
	//or nil if all jobs are done
	errCh chan error
	//wg is used to wait until all jobs are done and all workers are stopped listening
	wg *sync.WaitGroup
}

func New(ctx context.Context, parallel uint, handler Handler) *Pool {
	pool := &Pool{
		ch:    make(chan string),
		errCh: make(chan error),
		wg:    &sync.WaitGroup{},
	}

	for i := uint(0); i < parallel; i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()

			w := &worker{
				Handler: handler,
				ch:      pool.ch,
			}

			if err := w.listen(ctx); err != nil {
				pool.errCh <- err
			}
		}()
	}

	return pool
}

func (p *Pool) Handle(urls ...string) error {
	go func() {
		for _, url := range urls {
			p.ch <- url
		}
		close(p.ch)
	}()

	go func() {
		//when all workers are stopped nil is sent to error channel
		//to notify end of handling
		p.wg.Wait()
		p.errCh <- nil
	}()

	//returning the first error that passed to error channel from
	//any worker or nil if there were no errors and all jobs are done
	return <-p.errCh
}
