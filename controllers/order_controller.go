package controllers

import (
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/middlewares"
	"simple-ecommerce/models"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
	"simple-ecommerce/types"
	"strconv"
	"time"
)

type OrderApiControllerInterface interface {
	List(c *fiber.Ctx) error
	ByCustomerWantList(c *fiber.Ctx) error
	ListHtml(c *fiber.Ctx) error
	ByCustomerWantDetailHtml(c *fiber.Ctx) error
	ByCustomerWantCompleteHtml(c *fiber.Ctx) error
	ByCustomerWantGet(c *fiber.Ctx) error
	ByCustomerWantSave(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type orderApiControllers struct{}

var (
	OrderApiController OrderApiControllerInterface = &orderApiControllers{}
)

func (ctl *orderApiControllers) List(c *fiber.Ctx) error {
	responseList := services.GetOrderService().GetAllWithProducts()
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Status(fiber.StatusOK).JSON(responseList)
}

func (ctl *orderApiControllers) ByCustomerWantList(c *fiber.Ctx) error {
	responseList := new(responses.ResponseList)
	claims := c.Locals("claims").(middlewares.MyCustomClaims)
	customerId := claims.UserID

	log.Info(claims)

	responseList = services.GetOrderService().GetAllByCustomerIdWithProducts(customerId)
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Status(fiber.StatusOK).JSON(responseList)
}

func (ctl *orderApiControllers) ListHtml(c *fiber.Ctx) error {

	responseList := services.GetOrderService().GetAllWithProducts()
	if !responseList.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(responseList)
	}
	return c.Render("order_list", fiber.Map{
		"Orders": responseList.Rows,
	})
}

func (ctl *orderApiControllers) ByCustomerWantDetailHtml(c *fiber.Ctx) error {
	resp := new(responses.Response)
	token := c.Params("token")
	resp = services.GetOrderService().GetByTokenWithProducts(token)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Render("order_reminder", fiber.Map{
		"Order":   resp.Data,
		"BaseURL": c.BaseURL(),
	})
}

func (ctl *orderApiControllers) ByCustomerWantCompleteHtml(c *fiber.Ctx) error {
	resp := new(responses.Response)
	token := c.Params("token")
	resp = services.GetOrderService().GetByTokenWithProducts(token)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	data, ok := resp.Data.(*models.Order)
	if !ok {
		log.Info(ok)
		resp.Message = "Internal Error"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	err := services.GetOrderService().StatusUpdate(data.Id, "Completed")
	if err != nil {
		log.Info(ok)
		resp.Message = "Internal Error"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	return c.Render("order_complete", fiber.Map{
		"Order":   resp.Data,
		"BaseURL": c.BaseURL(),
	})
}

func (ctl *orderApiControllers) Get(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetOrderService().GetByIdWithProducts(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ctl *orderApiControllers) ByCustomerWantGet(c *fiber.Ctx) error {
	resp := new(responses.Response)
	claims := c.Locals("claims").(middlewares.MyCustomClaims)
	customerId := claims.UserID

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetOrderService().GetByIdAndCustomerIdWithProducts(id, customerId)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ctl *orderApiControllers) ByCustomerWantSave(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false
	claims := c.Locals("claims").(middlewares.MyCustomClaims)
	customerId := claims.UserID
	req := requests.OrderSaveRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		req.CustomerId = customerId
		req.OrderDate = types.JsonDate(time.Now())
		req.Status = "Pending"
		log.Info(req.OrderDate.ToTimeString())
		resp = services.GetOrderService().Save(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *orderApiControllers) Save(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.OrderSaveRequest{}
	if err := c.BodyParser(&req); err != nil {
		resp.Message = "Bad Request : Invalid JSON body"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	} else {
		resp = services.GetOrderService().Save(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *orderApiControllers) Update(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.OrderUpdateRequest{}
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
		resp = services.GetOrderService().Update(&req)
		if !resp.Success {
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

func (ctl *orderApiControllers) Delete(c *fiber.Ctx) error {
	resp := new(responses.Response)
	resp.Success = false

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		resp.Message = "Bad Request : Bad Request"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = services.GetOrderService().DeleteById(id)
	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
