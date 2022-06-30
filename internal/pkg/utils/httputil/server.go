package httputil

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupGracefulStop(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("Shutdown Server ...")
	Shutdown(srv)
}

func Shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("[http server shutdown err:]", err)
	}

	select {
	case <-ctx.Done():
		logrus.Println("[http server exit timeout of 5 seconds.]")
	default:

	}
	logrus.Printf("[http server exited.]")
}
