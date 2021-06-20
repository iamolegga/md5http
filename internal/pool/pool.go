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
	//count of workers
	parallel uint
	//handler that will do job, will be passed to each worker
	handler Handler
}

func New(parallel uint, handler Handler) *Pool {
	return &Pool{
		ch:       make(chan string),
		errCh:    make(chan error),
		wg:       &sync.WaitGroup{},
		parallel: parallel,
		handler:  handler,
	}
}

func (p *Pool) Handle(ctx context.Context, urls ...string) error {
	p.startWorkers(ctx)
	return p.process(urls)
}

func (p *Pool) startWorkers(ctx context.Context) {
	for i := uint(0); i < p.parallel; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()

			w := &worker{
				Handler: p.handler,
				ch:      p.ch,
			}

			if err := w.listen(ctx); err != nil {
				p.errCh <- err
			}
		}()
	}
}

func (p *Pool) process(urls []string) error {
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
