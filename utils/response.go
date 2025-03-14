package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Remark string      `json:"remark"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, remark string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Status: "success",
		Code:   statusCode,
		Remark: remark,
		Data:   data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, remark string) error {
	return c.Status(statusCode).JSON(Response{
		Status: "error",
		Code:   statusCode,
		Remark: remark,
	})
}
