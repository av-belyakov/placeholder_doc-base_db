package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)

	go func() {
		<-ctx.Done()

		fmt.Println("Placeholder_doc-base_db module is stop")

		stop()
	}()

	server(ctx)
}
