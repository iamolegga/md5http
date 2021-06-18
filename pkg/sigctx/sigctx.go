package sigctx

import (
	"context"
	"os"
	"os/signal"
)

//New returns context that will be canceled on receiving passed signal
func New(parent context.Context, signals ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(parent)

	go func() {
		c := make(chan os.Signal)
		defer close(c)

		signal.Notify(c, signals...)
		<-c
		signal.Stop(c)

		cancel()
	}()

	return ctx
}
