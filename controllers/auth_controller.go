package controllers

import (
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
)

type AuthApiControllerInterface interface {
	LoginCustomer(c *fiber.Ctx) error
	LoginAdmin(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	WelcomeCustomer(c *fiber.Ctx) error
	WelcomeAdmin(c *fiber.Ctx) error
	WelcomeGuest(c *fiber.Ctx) error
}

type authApiontrollers struct{}

var (
	AuthApiController AuthApiControllerInterface = &authApiontrollers{}
)

func (ctl *authApiontrollers) LoginCustomer(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.AuthLoginCustomerRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		log.Info(req)
		resp = services.GetCustomerService().Login(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *authApiontrollers) LoginAdmin(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.AuthLoginAdminRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		log.Info(req)
		resp = services.GetAdminService().Login(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *authApiontrollers) Register(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.CustomerRegisterRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		resp = services.GetCustomerService().Register(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *authApiontrollers) WelcomeCustomer(c *fiber.Ctx) error {
	return c.SendString("Welcome Customer!")
}

func (ctl *authApiontrollers) WelcomeAdmin(c *fiber.Ctx) error {
	return c.SendString("Welcome Admin!")
}

func (ctl *authApiontrollers) WelcomeGuest(c *fiber.Ctx) error {
	return c.SendString("Guest!")
}
