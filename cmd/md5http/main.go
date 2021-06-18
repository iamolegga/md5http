package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/iamolegga/md5http/internal/config"
	"github.com/iamolegga/md5http/internal/fetcher"
	"github.com/iamolegga/md5http/internal/handler"
	"github.com/iamolegga/md5http/internal/md5hasher"
	"github.com/iamolegga/md5http/internal/pool"
	"github.com/iamolegga/md5http/pkg/sigctx"
)

func main() {
	ctx := sigctx.New(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGABRT,
		syscall.SIGTERM,
	)

	cfg := config.New()

	h := handler.New(
		fetcher.New(),
		md5hasher.New(),
		os.Stdout,
	)

	p := pool.New(ctx, cfg.Parallel, h)

	switch err := p.Handle(cfg.Inputs...); err {
	case nil:
		return
	case context.Canceled:
		fmt.Println("\ncanceled")
	default:
		fmt.Printf("unexpected error: %v\n", err)
		os.Exit(1)
	}
}
