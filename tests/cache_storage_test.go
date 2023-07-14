package tests

import (
	. "cache-ahead-task/cache"
	mock_db "cache-ahead-task/db/mocks"
	. "cache-ahead-task/logger"
	"cache-ahead-task/model"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	Logger = zap.NewNop()
	m.Run()
}

func TestSliceBorders(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	ProductsSource := mock_db.NewMockProductsSource(c)
	products := []model.Product{
		{Id: "test1", Price: 1},
		{Id: "test2", Price: 2},
		{Id: "test3", Price: 3},
		{Id: "test4", Price: 4},
		{Id: "test5", Price: 5},
	}
	ProductsSource.EXPECT().GetProducts().Return(products)
	CommonCache.UpdateProducts(ProductsSource)

	t.Run("ordinary", func(t *testing.T) {
		result := CommonCache.GetProductsSlice(3, 1)
		if NotEqual(result, products[1:4]) {
			t.Errorf("%v != %v", result, products[1:4])
		}
	})

	t.Run("negative limit", func(t *testing.T) {
		result := CommonCache.GetProductsSlice(-3, 0)
		if NotEqual(result, products) {
			t.Errorf("%v != %v", result, products)
		}
	})

	t.Run("negative offset", func(t *testing.T) {
		result := CommonCache.GetProductsSlice(3, -1)
		if NotEqual(result, products[0:3]) {
			t.Errorf("%v != %v", result, products[0:3])
		}
	})

	t.Run("left border out of bounds", func(t *testing.T) {
		result := CommonCache.GetProductsSlice(10, 4)
		if NotEqual(result, products[4:]) {
			t.Errorf("%v != %v", result, products[4:])
		}
	})

	t.Run("right border out of bounds", func(t *testing.T) {
		result := CommonCache.GetProductsSlice(1, 10)
		if len(result) != 0 {
			t.Errorf("%v != %v", result, []model.Product{})
		}
	})
}

func TestUpdatingCache(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	ProductsSource := mock_db.NewMockProductsSource(c)
	products1 := []model.Product{{Id: "test", Price: 1}}
	products2 := []model.Product{{Id: "test", Price: 2}}
	duration := time.Millisecond
	gomock.InOrder(
		ProductsSource.EXPECT().GetProducts().Return(products1),
		ProductsSource.EXPECT().GetProducts().Return(products2).AnyTimes(),
	)

	CommonCache.GoPeriodicProductsUpdates(duration, ProductsSource)
	time.Sleep(duration)
	result1 := CommonCache.GetProductsSlice(1, 0)
	time.Sleep(duration)
	result2 := CommonCache.GetProductsSlice(1, 0)

	if result1[0] != products1[0] || result2[0] != products2[0] {
		t.Errorf("something is wrong with periodic cache update: cache1=%v, data1=%v, cache2=%v, data2=%v",
			result1, products1, result2, products2)
	}
}

func TestCacheSecurity(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	ProductsSource := mock_db.NewMockProductsSource(c)
	products := []model.Product{{Id: "test", Price: 1}}
	ProductsSource.EXPECT().GetProducts().Return(products)

	CommonCache.UpdateProducts(ProductsSource)
	result1 := CommonCache.GetProductsSlice(1, 0)
	result1[0].Price++
	result2 := CommonCache.GetProductsSlice(1, 0)

	if result2[0].Price != products[0].Price {
		t.Error("cache data can be changed by the received link")
	}
}

func NotEqual(a, b []model.Product) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}
