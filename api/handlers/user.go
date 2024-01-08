package handlers

import (
	"fabiloco/hotel-trivoli-api/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    users, error := service.FetchUsers()

    if error != nil {
      ctx.Locals("data", fiber.Map{
        "errors": error.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }

    ctx.Locals("data", fiber.Map{
      "users": users,
    })

    return ctx.SendStatus(fiber.StatusOK)
  }
}

func GetUserById(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")
    if err != nil {
      ctx.Locals("data", fiber.Map{
        "errors": err.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }

    user, error := service.FetchUserById(uint(id))

    if error != nil {
      ctx.Locals("data", fiber.Map{
        "errors": error.Error(),
      })
      return ctx.SendStatus(fiber.StatusNotFound)
    }

    ctx.Locals("data", fiber.Map{
      "user": user,
    })

    return ctx.SendStatus(fiber.StatusOK)
  }
}

func PostUsers(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
  //
  // var body entities.CreateUser
  // 
  // if err := ctx.BodyParser(&body); err != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": err.Error(),
		// })
  //   return ctx.SendStatus(fiber.StatusBadRequest)
  // }
  //
  // validationErrors := utils.ValidateInput(ctx, body)
  //
  // if validationErrors != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": validationErrors,
		// })
  //   return ctx.SendStatus(fiber.StatusBadRequest)
  // }
  //
  // var productTypesSlice []entities.UserType
  //
  // for i := 0; i < len(body.Type); i++{
  //   productType, error := h.productTypeStore.FindById(body.Type[i])
  //
  //   if error != nil {
  //     ctx.Locals("data", fiber.Map{
  //       "errors": error.Error() + fmt.Sprintf(": product type with id: %d", body.Type[i]),
  //     })
  //     return ctx.SendStatus(fiber.StatusNotFound)
  //   }
  //
  //   productTypesSlice = append(productTypesSlice, *productType)
  // }
  //
  // newUser := model.User {
  //   Name: body.Name,
  //   Price: body.Price,
  //   Stock: body.Stock,
  //   Type: productTypesSlice,
  // }
  //
  // product, error := h.productStore.Create(&newUser)
  //
  // if error != nil {
  //   ctx.Locals("data", fiber.Map{
		// 	"errors": error.Error(),
		// })
  //   return ctx.SendStatus(fiber.StatusNotFound)
  // }
  //
  // ctx.Locals("data", fiber.Map{
  //   "product": product,
  // })
  return ctx.SendStatus(fiber.StatusCreated)
  }
}

func PutUser(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    // var body entities.CreateUser
    //
    // if err := ctx.BodyParser(&body); err != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": err.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // validationErrors := utils.ValidateInput(ctx, body)
    //
    // if validationErrors != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": validationErrors,
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    // 
    // id, err := ctx.ParamsInt("id")
    //
    // if err != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": err.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // var productTypesSlice []model.UserType
    //
    // for i := 0; i < len(body.Type); i++{
    //   productType, error := h.productTypeStore.FindById(body.Type[i])
    //
    //   if error != nil {
    //     ctx.Locals("data", fiber.Map{
    //       "errors": error.Error() + fmt.Sprintf(": product type with id: %d", body.Type[i]),
    //     })
    //     return ctx.SendStatus(fiber.StatusNotFound)
    //   }
    //
    //   productTypesSlice = append(productTypesSlice, *productType)
    // }
    //
    // newUser := model.User {
    //   Name: body.Name,
    //   Price: body.Price,
    //   Stock: body.Stock,
    //   Type: productTypesSlice,
    // }
    //
    //
    // product, error := h.productStore.Update(uint(id), &newUser)
    //
    // if error != nil {
    //   ctx.Locals("data", fiber.Map{
    //     "errors": error.Error(),
    //   })
    //   return ctx.SendStatus(fiber.StatusBadRequest)
    // }
    //
    // ctx.Locals("data", fiber.Map{
    //   "product": product,
    // })
    //
    return ctx.SendStatus(fiber.StatusCreated)
  }
}

func DeleteUserById(service user.Service) fiber.Handler {
  return func(ctx *fiber.Ctx) error {
    id, err := ctx.ParamsInt("id")

    if err != nil {
      ctx.Locals("data", fiber.Map{
        "errors": err.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }
    
    user, error := service.RemoveUser(uint(id))

    if error != nil {
      ctx.Locals("data", fiber.Map{
        "errors": error.Error(),
      })
      return ctx.SendStatus(fiber.StatusBadRequest)
    }

    ctx.Locals("data", fiber.Map{
      "user": user,
    })

    return ctx.SendStatus(fiber.StatusOK)
  }
}

