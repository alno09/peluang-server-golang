package category

import (
	"peluang-server/domain"
	"peluang-server/dto"
	"peluang-server/internal/middleware"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type route struct {
	cateService domain.CategoryService
}

func newRoute(app *fiber.App, cateService domain.CategoryService) {
	route := route{
		cateService,
	}

	protectedApi := app.Group("/api")
	protectedApi.Use(middleware.Authenticate())
	{
		protectedApi.Get("/category", route.GetCategory)
	}
}

func (r *route) GetCategory(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(&dto.HttpResponse{
			Message: "Unauthorized",
			Data:    []string{},
		})
	}

	category, err := r.cateService.GetAllCAtegory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&dto.HttpResponse{
			Message: "Error",
			Data:    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.HttpResponse{
		Message: "success",
		Data:    category,
	})
}
