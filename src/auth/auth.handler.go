package Auth

import (
	"FM/src/configuration"
	"FM/src/core/exception"
	"FM/src/core/http"
	"FM/src/core/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService
	configuration.Config
}

func NewAuthHandler(authService *AuthService, config configuration.Config) *AuthHandler {
	return &AuthHandler{AuthService: *authService, Config: config}
}

func (handler AuthHandler) Route(app *fiber.App) {
	var basePath = utils.GetBaseRoute(handler.Config, "/auth")
	route := app.Group(basePath)

	route.Post("/login", handler.SignInWithGoogle)
}

func (handler AuthHandler) SignInWithGoogle(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(http.HttpResponse{
			StatusCode: fiber.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
	}

	var requestData struct {
		IDToken string `json:"idToken"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		exception.HandleError(c, err)
	}

	idToken := requestData.IDToken

	result, err := handler.AuthService.SignInWithGoogle(c.Context(), idToken)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    "Sign in with google failed",
			Data:       err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(http.HttpResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Sign in with google successfully",
		Data:       result,
	})
}
