package routers

import (
	"github.com/an-0305/pinot-client/handlers"
	"github.com/an-0305/pinot-client/repositories"
	"github.com/an-0305/pinot-client/services"
	"github.com/labstack/echo/v4"
	"github.com/startreedata/pinot-client-go/pinot"
)

type Repositories struct {
	salesRepo repositories.SalesRepository
}

type Services struct {
	salesSvc services.SalesService
}

type Handlers struct {
	salesHdlr handlers.SalesHandler
}

func Init(e *echo.Echo, db *pinot.Connection) {
	repo := &Repositories{
		salesRepo: repositories.NewSalesRepository(db),
	}
	svcs := &Services{
		salesSvc: services.NewSalesService(repo.salesRepo),
	}
	hdlrs := &Handlers{
		salesHdlr: handlers.NewSalesHandler(svcs.salesSvc),
	}

	e.GET("/sales/summaries", hdlrs.salesHdlr.GetSalesSummaryGroupByStoreAndDivision)
	e.GET("/sales/summaries/:storeCode", hdlrs.salesHdlr.GetSalesSummaryByStoreAndDivision)
}
