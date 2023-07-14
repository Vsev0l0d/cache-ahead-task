package handler

import (
	. "cache-ahead-task/cache"
	"cache-ahead-task/config"
	. "cache-ahead-task/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 {
		limit = config.DefaultLimit
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = config.DefaultOffset
	}
	Logger.Info(fmt.Sprintf("request with {limit=%d, offset=%d} (validated)", limit, offset))

	products := CommonCache.GetProductsSlice(limit, offset)
	responseJson, _ := json.Marshal(products)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(responseJson)
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}
