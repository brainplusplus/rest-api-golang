package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
	"strconv"
	"time"
)

type ProductApiControllerInterface interface {
	List(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type productApiControllers struct{}

var (
	ProductApiController ProductApiControllerInterface = &productApiControllers{}
)

func (ctl *productApiControllers) List(c *fiber.Ctx) error {

	responseList := services.GetProductService().GetAll()
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Status(fiber.StatusOK).JSON(responseList)
}

func (ctl *productApiControllers) Get(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetProductService().GetById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ctl *productApiControllers) Save(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.ProductSaveRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		file, err := c.FormFile("image")
		hasFile := true
		if err != nil {
			log.Info(err)
			hasFile = false
		}
		fileName := ""
		if hasFile {
			fileName = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file.Filename)
			err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileName))
			if err != nil {
				log.Info(err)
				resp.Message = "Failed to Upload Image"
				return c.Status(fiber.StatusInternalServerError).JSON(resp)
			}
		}
		req.Image = fileName
		resp = services.GetProductService().Save(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *productApiControllers) Update(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.ProductUpdateRequest{}
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
		file, err := c.FormFile("image")
		hasFile := true
		if err != nil {
			log.Info(err)
			hasFile = false
		}
		fileName := ""
		if hasFile {
			fileName = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file.Filename)
			err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileName))
			if err != nil {
				log.Info(err)
				resp.Message = "Failed to Upload Image"
				return c.Status(fiber.StatusInternalServerError).JSON(resp)
			}
		}
		req.Image = fileName
		resp = services.GetProductService().Update(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *productApiControllers) Delete(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request : Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetProductService().DeleteById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
