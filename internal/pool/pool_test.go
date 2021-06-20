package pool_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/iamolegga/md5http/internal/pool"
	"github.com/iamolegga/md5http/internal/pool/mock"
)

func TestPool_HandlePositive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	h := mock.NewMockHandler(mockCtrl)

	url := "https://google.com"
	h.EXPECT().Handle(url).Return(nil).Times(1)

	p := pool.New(1, h)
	if err := p.Handle(context.Background(), url); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPool_HandleNegative(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	h := mock.NewMockHandler(mockCtrl)

	url := "https://google.com"
	errHandling := errors.New("handling error")
	h.EXPECT().Handle(url).Return(errHandling).Times(1)

	p := pool.New(1, h)
	if err := p.Handle(context.Background(), url); !errors.Is(err, errHandling) {
		t.Errorf("error want: %v, got: %v", errHandling, err)
	}
}

func TestPool_HandleParallelPositive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	h := mock.NewMockHandler(mockCtrl)

	urls := []string{"https://google.com", "https://bing.com", "https://yahoo.com"}

	h.EXPECT().Handle(gomock.AssignableToTypeOf(urls[0])).DoAndReturn(func(_ string) error {
		time.Sleep(time.Second)
		return nil
	}).Times(3)

	p := pool.New(uint(len(urls)), h)

	start := time.Now()
	if err := p.Handle(context.Background(), urls...); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	elapsedSec := time.Since(start).Seconds()

	if elapsedSec < 1 || elapsedSec >= 2 {
		t.Errorf("enexpected time elapsed: %f", elapsedSec)
	}
}

func TestPool_HandleParallelNegative(t *testing.T) {

	//in this test, we are using 2 workers for 9 jobs.
	//let's force each job to take N sec, where N is idx + 1
	//and let 6th job return error, so handling should be like this:
	//		sec	0	1	2	3	4	5	6	7	8
	//	worker
	//	1		①	③	③	③	⑤	⑤	⑤	⑤	⑤
	//	2		②	②	④	④	④	④	⑥	-	-
	//number in circle is a job number that holds certain worker in certain second
	//expecting 3 things:
	//- returning error from 6th job
	//- handler should be called 6 times
	//- elapsed time should be less then 7 seconds (2 + 4 + 0.0xxx)

	mockCtrl := gomock.NewController(t)
	h := mock.NewMockHandler(mockCtrl)

	urls := []string{
		"https://google.ru", "https://bing.ru", "https://yahoo.ru",
		"https://google.com", "https://bing.com", "https://yahoo.com",
		"https://google.co.uk", "https://bing.co.uk", "https://yahoo.co.uk",
	}

	errHandling := errors.New("handling error")
	i := 0
	h.EXPECT().Handle(gomock.AssignableToTypeOf(urls[0])).DoAndReturn(func(_ string) error {
		i++
		if i == 6 {
			return errHandling
		}

		time.Sleep(time.Second * time.Duration(i))
		return nil
	}).Times(6)

	p := pool.New(2, h)

	start := time.Now()
	if err := p.Handle(context.Background(), urls...); !errors.Is(err, errHandling) {
		t.Errorf("unexpected error: %v", err)
	}
	elapsedSec := time.Since(start).Seconds()

	if elapsedSec < 6 || elapsedSec >= 7 {
		t.Errorf("enexpected time elapsed: %f", elapsedSec)
	}
}
