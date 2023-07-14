package main

import (
	. "cache-ahead-task/internal/cache"
	"cache-ahead-task/internal/config"
	db2 "cache-ahead-task/internal/db"
	"cache-ahead-task/internal/handler"
	. "cache-ahead-task/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	CommonCache.GoPeriodicProductsUpdates(config.DurationOfTheUpdatePeriod, db2.NewProductsSource(db2.Connect()))
	r := chi.NewRouter()
	r.Get("/", handler.Handler)
	Logger.Info("Listen and serve address http://" + config.HttpAddress)
	err := http.ListenAndServe(config.HttpAddress, r)
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}
