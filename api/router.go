package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/Sirupsen/logrus"
	"github.com/jchprj/GeoOrderTest/api/handlers"
	"github.com/jchprj/GeoOrderTest/cfg"
)

//Init init API http server
func Init() {
	var wait = cfg.HTTPServer.ShutdownTimeout

	r := handlers.GetRouter()

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Addr,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		Handler:      r,
	}
	go func() {
		logrus.WithFields(logrus.Fields{
			"Addr": cfg.HTTPServer.Addr,
		}).Infoln("Start API http server")
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatal("failed to start stress server", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// catch SIGINT (Ctrl+C), SIGKILL
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	logrus.Warn("preparing shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	logrus.Warn("shutting down")
	os.Exit(0)
}
