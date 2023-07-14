package main

import (
	. "cache-ahead-task/cache"
	"cache-ahead-task/config"
	"cache-ahead-task/db"
	"cache-ahead-task/handler"
	. "cache-ahead-task/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	CommonCache.GoPeriodicProductsUpdates(config.DurationOfTheUpdatePeriod, db.NewProductsSource(db.Connect()))
	r := chi.NewRouter()
	r.Get("/", handler.Handler)
	Logger.Info("Listen and serve address http://" + config.HttpAddress)
	err := http.ListenAndServe(config.HttpAddress, r)
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}
