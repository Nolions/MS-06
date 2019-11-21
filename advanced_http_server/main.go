package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=======start=======")

	var wg sync.WaitGroup
	wg.Add(3)
	ctx, cancel := context.WithCancel(context.Background())
	eg, egCtx := errgroup.WithContext(context.Background())

	eg.Go(helloWordServer(ctx, &wg))
	eg.Go(hellNameServer(ctx, &wg))
	eg.Go(echoServer(ctx, &wg))

	// server本身發生錯誤時終止所有服務
	go func() {
		<-egCtx.Done()
		cancel()
	}()

	// 收到系統發出SIGINT or SIGTERM通知時終止所有服務
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	if err := eg.Wait(); err != nil {
		fmt.Printf("error in the server goroutines: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("everything closed successfully")

	fmt.Println("=======end=======")
}

func helloWordServer(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {
		r := gin.Default()

		r.GET("/", helloWordHandler())

		s := &http.Server{
			Addr:              ":7000",
			Handler:           r,
			TLSConfig:         nil,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       0 * time.Second,
			MaxHeaderBytes:    1 << 20,
		}

		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := s.Shutdown(shutCtx); err != nil {
				errChan <- fmt.Errorf("error shutting down thing down the hello world servererror shutting down thing down the hello world server: %s\n", err)
				//fmt.Printf("error shutting down the hello world server: %s\n", err)
			}

			fmt.Println("the hello world server is closed")
			wg.Done()
			close(errChan)
		}()

		fmt.Println("the hello world server is starting")

		if err := s.ListenAndServe(); err != nil {
			//fmt.Printf("error starting the hello world server: %s\n", err)
			return fmt.Errorf("error starting the hello world server: %s\n", err)
		}

		fmt.Println("the hello world server is closing")
		err := <-errChan
		wg.Wait()

		return err
	}
}

func hellNameServer(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {
		r := gin.Default()
		r.GET("/", helloNameHandler())

		s := &http.Server{
			Addr:              ":7001",
			Handler:           r,
			TLSConfig:         nil,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       0 * time.Second,
			MaxHeaderBytes:    1 << 20,
		}

		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := s.Shutdown(shutCtx); err != nil {
				//fmt.Printf("error shutting down the hello name server: %s\n", err)
				errChan <- fmt.Errorf("error shutting down the hello name server: %s\n", err)
			}

			fmt.Println("the hello world server is closed")
			close(errChan)
			wg.Done()
		}()

		fmt.Println("the hello name server is starting")
		if err := s.ListenAndServe(); err != nil {
			//fmt.Printf("error starting the hello name server: %s\n", err)
			return fmt.Errorf("error starting the hello name server: %s\n", err)
		}

		fmt.Println("the hello name server is closing")
		err := <-errChan
		wg.Wait()
		return err
	}
}

func echoServer(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {

		r := gin.Default()
		r.POST("/", echoHandler())

		s := &http.Server{
			Addr:              ":7002",
			Handler:           r,
			TLSConfig:         nil,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       0 * time.Second,
			MaxHeaderBytes:    1 << 20,
		}

		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := s.Shutdown(shutCxt); err != nil {
				//fmt.Printf("error shutting down the echo server: %s\n", err)
				errChan <- fmt.Errorf("error shutting down the echo server: %s\n", err)
			}

			fmt.Println("the echo server is closing")
			close(errChan)
			wg.Done()
		}()

		fmt.Println("the echo server is starting")
		if err := s.ListenAndServe(); err != nil {
			//fmt.Printf("error starting the echo server: %s\n", err)
			return fmt.Errorf("error starting the echo server: %s\n", err)
		}

		fmt.Println("the echo server is closing")
		err := <-errChan
		wg.Wait()
		return err
	}
}

func helloWordHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusSeeOther, "HelloWord")
	}
}

func helloNameHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		n := context.Query("name")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s", n))
	}
}

func echoHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		body, _ := ioutil.ReadAll(context.Request.Body)
		context.String(http.StatusOK, string(body))
	}
}
