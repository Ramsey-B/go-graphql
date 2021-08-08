package mutations

import (
	"context"
	"go-graphql/graph/model"

	"github.com/jinzhu/gorm"
)

type OrderMutation struct {
	DB *gorm.DB
}

func (m *OrderMutation) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}
	err := m.DB.Create(&order).Error
	return &order, err
}

func (m *OrderMutation) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	updatedOrder := model.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}
	err := m.DB.Save(&updatedOrder).Error
	return &updatedOrder, err
}

func (m *OrderMutation) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	err := m.DB.Where("order_id = ?", orderID).Delete(&model.Item{}).Error
	if err != nil {
		return false, err
	}
	err = m.DB.Where("order_id = ?", orderID).Delete(&model.Order{}).Error
	return err != nil, err
}

func mapItemsFromInput(itemsInput []*model.ItemInput) []*model.Item {
	var items []*model.Item
	for _, itemInput := range itemsInput {
		items = append(items, &model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
		})
	}
	return items
}
