package controllers

import (
	"github.com/gofiber/fiber/v2"

	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
	"strconv"
)

type AdminApiControllerInterface interface {
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type adminApiControllers struct{}

var (
	AdminApiController AdminApiControllerInterface = &adminApiControllers{}
)

func (ctl *adminApiControllers) List(c *fiber.Ctx) error {

	responseList := services.GetAdminService().GetAll()
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Status(fiber.StatusOK).JSON(responseList)
}

func (ctl *adminApiControllers) Get(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetAdminService().GetById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ctl *adminApiControllers) Save(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.AdminSaveRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		resp = services.GetAdminService().Save(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *adminApiControllers) Update(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.AdminUpdateRequest{}
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
		resp = services.GetAdminService().Update(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *adminApiControllers) Delete(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request : Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetAdminService().DeleteById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
