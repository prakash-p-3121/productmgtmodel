package productmgtmodel

import (
	"sync"
	"time"
)

type Category struct {
	ID         string    `json:"id"`
	IDBitCount uint64    `json:"id-bit-count"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created-at"`
	UpdatedAt  time.Time `json:"updated-at"`
}

type Product struct {
	ID          string    `json:"id"`
	IDBitCount  uint64    `json:"id-bit-count"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Currency    string    `json:"currency"`
	CostPrice   float64   `json:"cost-price"`
	CategoryID  string    `json:"category-id"`
	ImagePath   string    `json:"image-path"`
	CreatedAt   time.Time `json:"created-at"`
	UpdatedAt   time.Time `json:"updated-at"`
}

const (
	Return           uint = 1
	Refund           uint = 2
	ReturnOrRefund   uint = 3
	NoReturnOrRefund uint = 4
)

var validReturnPolicies sync.Map

func init() {
	validReturnPolicies.Store(Return, struct{}{})
	validReturnPolicies.Store(Refund, struct{}{})
	validReturnPolicies.Store(ReturnOrRefund, struct{}{})
	validReturnPolicies.Store(NoReturnOrRefund, struct{}{})
}

type ProductSellerAssociations struct {
	ID           string    `json:"id"`
	IDBitCount   uint64    `json:"id-bit-count"`
	ProductID    string    `json:"product-id"`
	SellerID     string    `json:"seller-id"` /* from userID table in usermgtms */
	SellingPrice float64   `json:"selling-price"`
	Currency     string    `json:"currency"`
	StockCount   *uint64   `json:"stock-count"`
	ReturnPolicy uint      `json:"return-policy"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}
