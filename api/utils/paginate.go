package utils

import (
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Paginate(c *fiber.Ctx, total int64, data interface{}) *entities.Pagination {
	// Leer params
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Calcular paginas
	pages := int((total + int64(limit) - 1) / int64(limit)) // ceil

	return &entities.Pagination{
		Page:  page,
		Limit: limit,
		Total: total,
		Pages: pages,
		Data:  data,
	}
}

func GetPaginationParams(c *fiber.Ctx) (page, limit, offset int) {
	page, _ = strconv.Atoi(c.Query("page", "1"))
	limit, _ = strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset = (page - 1) * limit
	return
}
