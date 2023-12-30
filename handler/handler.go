package handler

import "fabiloco/hotel-trivoli-api/product"

type Handler struct {
	productStore product.Store
}

func NewHandler(productStore product.Store) *Handler {
	return &Handler{
    productStore: productStore,
	}
}
