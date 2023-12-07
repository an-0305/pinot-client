package services

import (
	"context"

	"github.com/an-0305/pinot-client/models"
	"github.com/an-0305/pinot-client/packages/util"
	"github.com/an-0305/pinot-client/repositories"
	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type SalesService interface {
	GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) ([]interface{}, error)
	GetSalesSummaryByStoreAndDivision(ctx context.Context, storeCode uint32) []interface{}
}

type salesService struct {
	salesRepo repositories.SalesRepository
}

func NewSalesService(salesRepo repositories.SalesRepository) *salesService {
	return &salesService{
		salesRepo: salesRepo,
	}
}

func (s *salesService) GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) ([]interface{}, error) {
	res, err := s.salesRepo.GetSalesSummaryGroupByStoreAndDivision(ctx)
	if err != nil {
		return make([]interface{}, 0), err
	}
	jsonRes := util.CreateJSONResponse(res, mapSalesBrokerResponse)

	return jsonRes, nil
}

func (s *salesService) GetSalesSummaryByStoreAndDivision(ctx context.Context, storeCode uint32) []interface{} {
	res, err := s.salesRepo.GetSalesSummaryByStoreAndDivision(ctx, storeCode)
	jsonRes := util.CreateJSONResponse(res, mapSalesBrokerResponse)
	if err != nil {
		log.Error(err)
	}
	return jsonRes
}

func mapSalesBrokerResponse(table *pinot.ResultTable, row []interface{}) interface{} {
	return models.SalesSummary{
		StoreCode:    util.GetInt32Value(table, row, "storeCode"),
		Store:        util.GetStringValue(table, row, "store"),
		DivisionCode: util.GetInt32Value(table, row, "divisionCode"),
		Division:     util.GetStringValue(table, row, "division"),
		TotalPrice:   util.GetInt32Value(table, row, "totalPrice"),
	}
}
