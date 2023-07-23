package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itkln/ease-backend/app/controllers"
	"github.com/itkln/ease-backend/app/models"
)

type Handler struct {
	ctrl *controllers.AuthControllerImpl
}

func NewHandler(ctrl *controllers.AuthControllerImpl) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Register(ctx *fiber.Ctx) error {
	return ctx.SendString("Register")
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	user := &models.UserAuth{}

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "parse data failed",
		})
	}

	if err := h.ctrl.LoginUser(user); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "The requested user not found",
		})
	}

	token, err := h.ctrl.GenerateJWT(user)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return ctx.SendString(token)
}

func (h *Handler) Profile(ctx *fiber.Ctx) error {
	return ctx.SendString("Profile")
}
