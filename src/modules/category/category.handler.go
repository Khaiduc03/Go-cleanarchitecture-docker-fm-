package category

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/middleware"
	"FM/src/core/shared"
	"FM/src/core/utils"
	modelCategory "FM/src/modules/category/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	CategoryService
	configuration.Config
}

func NewCategoryHandler(categoryService *CategoryService, config configuration.Config) *CategoryHandler {
	return &CategoryHandler{CategoryService: *categoryService, Config: config}
}

func (handler CategoryHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/category")

	route := app.Group(basePath, middleware.AuthMiddleware(handler.Config), middleware.RoleMiddleware([]string{"TEACHER"}))
	route.Get("/type", handler.FindAllByType)
	route.Get("/", handler.FindAll)
	route.Get("/:id", handler.FindById)
	route.Post("/", handler.Create)
	route.Put("/", handler.Update)
	route.Delete("/:id", handler.Delete)

}

func (handler CategoryHandler) FindAll(c *fiber.Ctx) error {
	categories, err := handler.CategoryService.FindAll(c.Context())
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get all category successfully",
		Data:       categories,
	})
}

func (handler CategoryHandler) FindById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	category, err := handler.CategoryService.FindById(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get category by id successfully",
		Data:       category,
	})
}

func (handler CategoryHandler) Create(c *fiber.Ctx) error {
	validator := shared.NewValidator()
	var request modelCategory.CreateCategoryReq
	if err := c.BodyParser(&request); err != nil {
		return exception.HandleError(c, err)
	}
	if err := validator.Validate(request); err != nil {
		return exception.HandleErrorCustomMessage(c, "Missing required fields")
	}


	message, err := handler.CategoryService.Create(c.Context(), request)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler CategoryHandler) Update(c *fiber.Ctx) error {

	validator := shared.NewValidator()

	var model modelCategory.UpdateCategoryReq
	if err := c.BodyParser(&model); err != nil {
		return exception.HandleError(c, err)
	}

	if err := validator.Validate(model); err != nil {
		return exception.HandleErrorCustomMessage(c, "Missing required fields")
	}

	message, err := handler.CategoryService.Update(c.Context(), model)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler CategoryHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return exception.HandleError(c, err)
	}

	message, err := handler.CategoryService.Delete(c.Context(), id)
	if err != nil {
		return exception.HandleError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    message,
	})
}

func (handler CategoryHandler) FindAllByType(c *fiber.Ctx) error {
	category_type := c.Query("type")

	categories, err := handler.CategoryService.FindAllCategoryByType(c.Context(), category_type)
	if err != nil {
		return exception.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Get all category by type successfully",
		Data:       categories,
	})
}
