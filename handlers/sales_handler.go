package handlers

import (
	"net/http"
	"strconv"

	"github.com/an-0305/pinot-client/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type SalesHandler interface {
	GetSalesSummaryGroupByStoreAndDivision(c echo.Context) error
	GetSalesSummaryByStoreAndDivision(c echo.Context) error
}

type salesHandler struct {
	salesSvc services.SalesService
}

func NewSalesHandler(salesSvc services.SalesService) *salesHandler {
	return &salesHandler{
		salesSvc: salesSvc,
	}
}

func (h *salesHandler) GetSalesSummaryGroupByStoreAndDivision(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.salesSvc.GetSalesSummaryGroupByStoreAndDivision(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (h *salesHandler) GetSalesSummaryByStoreAndDivision(c echo.Context) error {
	ctx := c.Request().Context()
	storeCode, err := strconv.Atoi(c.Param("storeCode"))
	if err != nil {
		log.Error(err)
	}
	res := h.salesSvc.GetSalesSummaryByStoreAndDivision(ctx, uint32(storeCode))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
