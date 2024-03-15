package productmgtmodel

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"strings"
)

type CategoryCreateReq struct {
	Name *string `json:"name"`
}

func (req *CategoryCreateReq) Validate() errorlib.AppError {
	if req.Name == nil {
		return errorlib.NewBadReqError("name-nil")
	}
	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return errorlib.NewBadReqError("name-empty")
	}
	return nil
}

type ProductCreateReq struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Currency    *string  `json:"currency"`
	CostPrice   *float64 `json:"cost-price"`
	CategoryID  *string  `json:"category-id"`
	MediaType   *uint    `json:"media-type"`
	MediaPath   *string  `json:"media-path"`
}

func (req *ProductCreateReq) Validate() errorlib.AppError {
	if req.Name == nil {
		return errorlib.NewBadReqError("name-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.Name) {
		return errorlib.NewBadReqError("name-empty")
	}

	if req.Description == nil {
		return errorlib.NewBadReqError("description-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.Description) {
		return errorlib.NewBadReqError("description-empty")
	}

	if req.Currency == nil {
		return errorlib.NewBadReqError("currency-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.Currency) {
		return errorlib.NewBadReqError("currency-empty")
	}

	if req.CostPrice == nil {
		return errorlib.NewBadReqError("cost-price-nil")
	}
	if *req.CostPrice <= 0 {
		return errorlib.NewBadReqError("cost-price<=0")
	}

	if req.CategoryID == nil {
		return errorlib.NewBadReqError("category-id-nil")
	}

	if req.MediaType == nil {
		return errorlib.NewBadReqError("media-path-nil")
	}

	if !(*req.MediaType == MediaTypeImage || *req.MediaType == MediaTypeVideo) {
		return errorlib.NewBadReqError("invalid-media-type")
	}

	if req.MediaPath == nil {
		return errorlib.NewBadReqError("media-path-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.MediaPath) {
		return errorlib.NewBadReqError("media-path-empty")
	}
	return nil
}

type MarketplaceListingCreateReq struct {
	ProductID    *string  `json:"product-id"`
	SellerID     *string  `json:"seller-id"` /* from userID table in usermgtms */
	SellingPrice *float64 `json:"selling-price"`
	Currency     *string  `json:"currency"`
	StockCount   *uint64  `json:"stock-count"`
	ReturnPolicy *uint    `json:"return-policy"`
}

func (req *MarketplaceListingCreateReq) Validate() errorlib.AppError {
	if req.ProductID == nil {
		return errorlib.NewBadReqError("product-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.ProductID) {
		return errorlib.NewBadReqError("product-id-empty")
	}

	if req.SellerID == nil {
		return errorlib.NewBadReqError("seller-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.SellerID) {
		return errorlib.NewBadReqError("seller-id-empty")
	}

	if req.SellingPrice == nil {
		return errorlib.NewBadReqError("selling-price-nil")
	}
	if *req.SellingPrice <= 0 {
		return errorlib.NewBadReqError("selling-price<=0")
	}

	if req.Currency == nil {
		return errorlib.NewBadReqError("currency-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.Currency) {
		return errorlib.NewBadReqError("currency-empty")
	}

	if req.ReturnPolicy == nil {
		return errorlib.NewBadReqError("return-policy-nil")
	}
	_, ok := validReturnPolicies.Load(*req.ReturnPolicy)
	if !ok {
		return errorlib.NewBadReqError("invalid-return-policy")
	}
	return nil
}
