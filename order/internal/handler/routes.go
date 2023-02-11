package handler

import "github.com/labstack/echo/v4"

func (h *OrderHandler) Register(v1 *echo.Group) {
	c := v1.Group("/order")
	c.POST("", h.Order)
}
