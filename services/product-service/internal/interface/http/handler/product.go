package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/product"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/sku"
	"github.com/yangpixi/GoMall/shared/response"
)

type ProductHandler struct {
	productHandler *product.CreateProductHandler
	skuHandler     *sku.CreateSKUHandler
}

type createProductRequest struct {
	Name        string `json:"name"`
	ShopID      int64  `json:"shopId,string"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type createSKURequest struct {
	ProductID     int64  `json:"productId,string"`
	Price         int    `json:"price"`
	Specification string `json:"specification"`
}

func NewProductHandler(productHandler *product.CreateProductHandler, skuHandler *sku.CreateSKUHandler) (*ProductHandler, error) {
	if productHandler == nil || skuHandler == nil {
		return nil, errors.New("invalid product handler")
	}

	return &ProductHandler{productHandler: productHandler, skuHandler: skuHandler}, nil
}

func (h *ProductHandler) CreateProductHandler(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.productHandler.Handle(c.Request.Context(), &product.CreateProductCommand{
		Name:        req.Name,
		ShopID:      req.ShopID,
		Description: req.Description,
		Status:      req.Status,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "product created successfully")

}

func (h *ProductHandler) CreateSKUHandler(c *gin.Context) {
	var req createSKURequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.skuHandler.Handle(c.Request.Context(), &sku.CreateSKUCommand{
		ProductID:     req.ProductID,
		Price:         req.Price,
		Specification: req.Specification,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "sku created successfully")
}
