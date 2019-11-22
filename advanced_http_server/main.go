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

	//eg.Go(helloWordServer(ctx, &wg))
	eg.Go(createHttpServer("hello word", ":7000", helloWordEngine(helloWordHandler()), ctx, &wg))
	eg.Go(createHttpServer("hello name", ":7001", helloNameEngine(helloNameHandler()), ctx, &wg))
	eg.Go(createHttpServer("echo", ":7002", echoNameEngine(echoHandler()), ctx, &wg))

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

func createHttpServer(name, addr string, engine *gin.Engine, ctx context.Context, wg *sync.WaitGroup, ) func() error {
	return func() error {
		s := run(engine, addr)

		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := s.Shutdown(shutCxt); err != nil {
				//fmt.Printf("error shutting down the echo server: %s\n", err)
				errChan <- fmt.Errorf("error shutting down the %s server: %s\n", name, err)
			}

			fmt.Printf("the %s server is closing\n", name)
			close(errChan)
			wg.Done()
		}()

		fmt.Printf("the %s server is starting\n", name)
		if err := s.ListenAndServe(); err != nil {
			//fmt.Printf("error starting the echo server: %s\n", err)
			return fmt.Errorf("error starting the %s server: %s\n", name, err)
		}

		fmt.Printf("the %s server is closing\n", name)
		err := <-errChan
		wg.Wait()
		return err
	}
}

func run(r *gin.Engine, addr string) *http.Server {
	s := &http.Server{
		Addr:              addr,
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       0 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	return s
}

func helloWordEngine(handler func(c *gin.Context), ) *gin.Engine {
	r := gin.Default()
	r.GET("/", handler)

	return r
}

func helloNameEngine(handler func(c *gin.Context), ) *gin.Engine {
	r := gin.Default()
	r.GET("/", handler)

	return r
}

func echoNameEngine(handler func(c *gin.Context), ) *gin.Engine {
	r := gin.Default()
	r.POST("/", handler)

	return r
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
