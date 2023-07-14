package cache

import (
	"cache-ahead-task/internal/config"
	"cache-ahead-task/internal/db"
	. "cache-ahead-task/internal/logger"
	"cache-ahead-task/internal/model"
	"fmt"
	"github.com/go-co-op/gocron"
	"sync"
	"time"
)

var CommonCache Cache

func init() {
	CommonCache = NewCache()
}

type Cache interface {
	GetProductsSlice(limit, offset int) []model.Product
	UpdateProducts(productsSource db.ProductsSource)
	GoPeriodicProductsUpdates(duration time.Duration, productsSource db.ProductsSource)
}

type cache struct {
	mu       sync.RWMutex
	products []model.Product
}

func NewCache() Cache {
	return &cache{}
}

func (c *cache) GetProductsSlice(limit, offset int) []model.Product {
	if limit < 0 {
		limit = config.DefaultLimit
	}
	if offset < 0 {
		offset = config.DefaultOffset
	}
	c.mu.RLock()
	index := min(offset+limit, len(c.products))
	offset = min(offset, len(c.products))
	result := make([]model.Product, index-offset)
	copy(result, c.products[offset:index])
	c.mu.RUnlock()
	Logger.Info(fmt.Sprintf("Products[%d,%d] data was read from cache", offset, index))
	return result
}

func (c *cache) UpdateProducts(productsSource db.ProductsSource) {
	updatedProducts := productsSource.GetProducts()
	result := make([]model.Product, len(updatedProducts))
	copy(result, updatedProducts)
	c.mu.Lock()
	c.products = result
	c.mu.Unlock()
	Logger.Info("Cached products data updated")
}

func (c *cache) GoPeriodicProductsUpdates(duration time.Duration, productsSource db.ProductsSource) {
	s := gocron.NewScheduler(time.UTC)
	job, err := s.Every(duration).Do(c.UpdateProducts, productsSource)
	if err != nil {
		Logger.Error(fmt.Sprintf("Job: %v, Error: %v", job, err))
		return
	}
	s.StartAsync()
	Logger.Info("Created a cache data update job with a period " + duration.String())
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
