package db

import (
	. "cache-ahead-task/logger"
	"cache-ahead-task/model"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=products_source.go -destination=mocks/mock_products_source.go

type ProductsSource interface {
	GetProducts() []model.Product
}

type productsSource struct {
	DB *sqlx.DB
}

func NewProductsSource(db *sqlx.DB) ProductsSource {
	return &productsSource{DB: db}
}

func (s *productsSource) GetProducts() []model.Product {
	var products []model.Product
	err := s.DB.Select(&products, "SELECT * FROM products")
	if err != nil {
		Logger.Error(err.Error())
		return products
	}
	Logger.Info("Products data received")
	return products
}
