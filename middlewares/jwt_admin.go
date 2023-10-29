package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func IsAdminMiddleware(c *fiber.Ctx) error {
	claims := c.Locals("claims").(MyCustomClaims)
	isAdmin := claims.Role == "Administrator"
	if isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Forbidden"})
	}
	return c.Next()
}
