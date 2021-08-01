package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	root, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(root)

	// server
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// start web server
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(ctx)
	})

	// deal signal
	g.Go(func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)

		select {
		case sig := <-ch:
			fmt.Printf("receive signal: %+v", sig)
			cancel()
			return nil
		case <-ctx.Done():
			fmt.Printf("receive global done")
			return ctx.Err()
		}
	})

	err := g.Wait()
	fmt.Printf("g wait error: %+v", err)
}
