package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	AddSrv       = ":8080"
	TemplatesDir = "./html/"
)

func main() {

	quit := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(quit, os.Interrupt)

	fileSrv := http.FileServer(http.Dir(TemplatesDir))

	server := &http.Server{
		Addr:    AddSrv,
		Handler: fileSrv,
	}

	go func() {
		<-quit
		log.Println("server quitting...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}

		close(done)
	}()

	log.Println("server starting on port 8080...")

	err := server.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), "closed") {
		log.Fatalf("error in server %s", err)
	}

	<-done
	log.Println("server exited.")

}
