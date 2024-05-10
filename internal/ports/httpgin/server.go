// Package httpgin provides utilities to run http server
package httpgin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sveboo/exchangerate/internal/app"
	"golang.org/x/sync/errgroup"
)

// The type Server in Go represents a server with a port and a Gin engine.
type Server struct {
	port string
	app  *gin.Engine
}

// loggerMW represents simple logger handler with useful info about request and time to request
func loggerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		log.Println("\n[\n\r", "Time:", latency, "\nMethod used:",
			c.Request.Method, "\nRequest to:", c.Request.URL.Path, "\nStatus code:", status, "\n]")
	}
}

func captureSigQuit(ctx context.Context) func() error {
	return func() error {
		sigOuit := make(chan os.Signal, 1)
		signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
		signal.Notify(sigOuit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case s := <-sigOuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v ", s)

		case <-ctx.Done():
			return nil
		}
	}
}

// NewHTTPServer creates a new HTTP server with specified port and application handler using the Gin framework.
func NewHTTPServer(port string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	api := router.Group("info")
	s := &http.Server{Addr: port, Handler: router}
	api.Use(loggerMW(), gin.Recovery())
	AppRouter(api, a)
	return s

}

func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.app
}

func Run(a app.App, httpPort string) error {
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(captureSigQuit(ctx))

	eg.Go(Serve(ctx, a, httpPort))
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// Serve returns function to start HTTP server on a port given and implements graceful shutdown principle
func Serve(ctx context.Context, a app.App, httpPort string) func() error {
	return func() error {
		httpServer := NewHTTPServer(httpPort, a)

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", httpServer.Addr, err)
			}

			close(errCh)
		}()

		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	}
}
