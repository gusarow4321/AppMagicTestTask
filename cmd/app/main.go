package main

import (
	"AppMagicTestTask/internal/data"
	"AppMagicTestTask/internal/server"
	"AppMagicTestTask/internal/service"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var httpAddr string
var historyUrl string
var defaultUrl = "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	flag.StringVar(&historyUrl, "url", defaultUrl, "URL of price history json file")
	flag.StringVar(&httpAddr, "http", ":3000", "Address to listen on")
}

func newApp() (*http.Server, error) {
	results, err := data.NewResults(historyUrl)
	if err != nil {
		return nil, err
	}
	srv := service.NewService(results)
	if err != nil {
		return nil, err
	}

	return server.NewServer(httpAddr, srv), nil
}

func main() {
	flag.Parse()

	app, err := newApp()
	if err != nil {
		log.Panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = app.ListenAndServe(); err != nil {
			log.Printf("Server Failed: %v", err)
		}
	}()

	log.Print("Server Started")

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = app.Shutdown(ctx); err != nil {
		log.Printf("Shutdown Failed: %v", err)
	}

	log.Print("Server Stopped")
}
