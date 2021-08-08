package queries

import (
	"context"
	"go-graphql/graph/model"

	"gorm.io/gorm"
)

type OrderQuery struct {
	DB *gorm.DB
}

func (q *OrderQuery) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	err := q.DB.Preload("Items").Find(&orders).Error

	return orders, err
}
