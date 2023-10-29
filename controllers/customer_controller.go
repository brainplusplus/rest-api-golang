package controllers

import (
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
	"strconv"
)

type CustomerApiControllerInterface interface {
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type customerApiControllers struct{}

var (
	CustomerApiController CustomerApiControllerInterface = &customerApiControllers{}
)

func (ctl *customerApiControllers) List(c *fiber.Ctx) error {

	responseList := services.GetCustomerService().GetAll()
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Status(fiber.StatusOK).JSON(responseList)
}

func (ctl *customerApiControllers) Get(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetCustomerService().GetById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ctl *customerApiControllers) Save(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.CustomerSaveRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		resp = services.GetCustomerService().Save(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *customerApiControllers) Update(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.CustomerUpdateRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			resp.Message = "Bad Request : Bad Request"
			return c.Status(fiber.StatusBadRequest).JSON(resp)
		}
		req.Id = id
		resp = services.GetCustomerService().Update(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *customerApiControllers) Delete(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request : Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetCustomerService().DeleteById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
