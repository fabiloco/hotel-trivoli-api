// handler/handler.go
package handler

import (
	"fabiloco/hotel-trivoli-api/product"
	productType "fabiloco/hotel-trivoli-api/productType"
	"fabiloco/hotel-trivoli-api/user"
)

type Handler struct {
	productStore product.Store
	userStore    user.UserStore
	productTypeStore   productType.Store
}

func NewHandler(productStore product.Store, userStore user.UserStore, productTypeStore productType.Store) *Handler {
	return &Handler{
		productStore: productStore,
		userStore:    userStore,
    productTypeStore: productTypeStore,
	}
}
