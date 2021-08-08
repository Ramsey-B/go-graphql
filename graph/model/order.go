package model

type Item struct {
	ID          int    `json:"id" gorm:"primaryKey;column:id"`
	ProductCode string `json:"productCode" gorm:"column:product_code"`
	ProductName string `json:"productName" gorm:"column:product_name"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
}

type Order struct {
	ID           int     `json:"id" gorm:"column:id"`
	CustomerName string  `json:"customerName" gorm:"column:customer_name"`
	OrderAmount  float64 `json:"orderAmount" gorm:"column:order_amount"`
	Items        []*Item `json:"items" gorm:"foreignKey:id"`
}
