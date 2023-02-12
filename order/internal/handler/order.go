package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	rsErr "mac-laren/order/internal/error"
	"mac-laren/pkg/logger"
	"mac-laren/pkg/redis"
)

type OrderRequest struct {
	Price int64  `json:"price" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type OrderResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *OrderHandler) Order(ctx echo.Context) error {

	var request OrderRequest
	if err := ctx.Bind(&request); err != nil {
		logger.Logger.Error().Err(err).Msg("error while binding body request")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.ErrEmptyBodyRequest.Error(),
		})
	}

	// todo: add html escape

	b, err := json.Marshal(request)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while marshalling order request")
		return ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.ErrServer.Error(),
		})
	}

	if _, err := h.RedisRepository.LPush(redis.NEW_ORDER, string(b)); err != nil {
		logger.Logger.Error().Err(err).Msg("error while pushing new order to queue")
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: rsErr.ErrServer.Error(),
		})

	}

	// todo: improve response
	return ctx.JSON(http.StatusCreated, OrderResponse{
		Result: "order submitted",
	})

}
