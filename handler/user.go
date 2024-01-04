package handler

import (
	"fabiloco/hotel-trivoli-api/model"
	"fabiloco/hotel-trivoli-api/store"
	"fabiloco/hotel-trivoli-api/utils"

	"github.com/gofiber/fiber/v2"
)

// UserHandler es responsable de manejar las operaciones relacionadas con usuarios
type UserHandler struct {
	userStore *store.UserStore
}

// NewUserHandler crea una nueva instancia de UserHandler
func NewUserHandler(userStore *store.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// ListUsers retorna la lista de usuarios
func (h *Handler) ListUsers(ctx *fiber.Ctx) error {
	users, err := h.userStore.List()

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	ctx.Locals("data", fiber.Map{
		"users": users,
	})

	return ctx.SendStatus(fiber.StatusOK)
}

// GetUserById retorna un usuario por su ID
func (h *Handler) GetUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.userStore.FindById(id)

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	ctx.Locals("data", fiber.Map{
		"user": user,
	})

	return ctx.SendStatus(fiber.StatusOK)
}

// CreateUser crea un nuevo usuario
func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	var body model.CreateUser

	if err := ctx.BodyParser(&body); err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	validationErrors := utils.ValidateInput(ctx, body)

	if validationErrors != nil {
		ctx.Locals("data", fiber.Map{
			"errors": validationErrors,
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.userStore.Create(&body)

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	ctx.Locals("data", fiber.Map{
		"user": user,
	})

	return ctx.SendStatus(fiber.StatusCreated)
}

// UpdateUser actualiza un usuario existente
func (h *Handler) UpdateUser(ctx *fiber.Ctx) error {
	var body model.CreateUser

	if err := ctx.BodyParser(&body); err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	validationErrors := utils.ValidateInput(ctx, body)

	if validationErrors != nil {
		ctx.Locals("data", fiber.Map{
			"errors": validationErrors,
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.userStore.Update(id, &body)

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	ctx.Locals("data", fiber.Map{
		"user": user,
	})

	return ctx.SendStatus(fiber.StatusCreated)
}

// DeleteUserById elimina un usuario por su ID
func (h *Handler) DeleteUserById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.userStore.Delete(id)

	if err != nil {
		ctx.Locals("data", fiber.Map{
			"errors": err.Error(),
		})
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	ctx.Locals("data", fiber.Map{
		"user": user,
	})

	return ctx.SendStatus(fiber.StatusOK)
}
