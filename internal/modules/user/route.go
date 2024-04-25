package user

import (
	"fmt"
	"peluang-server/domain"
	"peluang-server/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type route struct {
	userService domain.UserService
}

func NewRoute(app *fiber.App, userService domain.UserService) {
	route := route{
		userService,
	}

	api := app.Group("/api")

	api.Post("/auth/register", route.UserRegister)
	api.Post("/auth/otp", route.ValidateOTP)
	api.Post("/auth/resend-otp", route.ResendOTP)
}
func (r *route) UserRegister(c *fiber.Ctx) error {
	user := new(dto.AuthRequest)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
		})
	}

	if err := validator.New().Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userModel := new(domain.User)
	userModel.Email = user.Email
	userModel.Password = user.Password
	userModel.Username = user.Username
	userModel.Telp = user.Telp

	otp, err := r.userService.Register(userModel, c.Context())
	if err != nil {
		if err == domain.ErrEmailExist {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": domain.ErrEmailExist.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"OTP": otp,
	})
}

func (r *route) ValidateOTP(c *fiber.Ctx) error {
	otp := new(dto.OTPRequest)
	if err := c.BodyParser(otp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("error parsing body: %v", err),
		})
	}

	if err := validator.New().Struct(otp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := r.userService.ValidateOTP(otp.UserID, otp.OTP)
	if err != nil {
		if err == domain.ErrInvalidOTP {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": domain.ErrInvalidOTP.Error(),
			})
		}
		if err == domain.ErrExpiredOTP {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": domain.ErrExpiredOTP.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("error validating otp: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "otp valid",
	})
}

func (r *route) ResendOTP(c *fiber.Ctx) error {
	otp := new(dto.ReOTPRequest)
	if err := c.BodyParser(otp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("error parsing body: %v", err),
		})
	}

	if err := validator.New().Struct(otp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newOtp, err := r.userService.ResendOTP(otp.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("error resending otp: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("otp %d has been sent", newOtp),
		"otp":     newOtp,
	})
}
