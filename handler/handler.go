// handler/handler.go
package handler

import (
	"fabiloco/hotel-trivoli-api/product"
	"fabiloco/hotel-trivoli-api/user"
)

type Handler struct {
	productStore product.Store
	userStore    user.UserStore
}

func NewHandler(productStore product.Store, userStore user.UserStore) *Handler {
	return &Handler{
		productStore: productStore,
		userStore:    userStore,
	}
}
