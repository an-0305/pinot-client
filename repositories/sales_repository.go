package repositories

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

type SalesRepository interface {
	GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) (*pinot.BrokerResponse, error)
	GetSalesSummaryByStoreAndDivision(ctx context.Context, storeCode uint32) (*pinot.BrokerResponse, error)
}

type salesRepository struct {
	db *pinot.Connection
}

func NewSalesRepository(db *pinot.Connection) *salesRepository {
	return &salesRepository{
		db: db,
	}
}

func (r *salesRepository) GetSalesSummaryGroupByStoreAndDivision(ctx context.Context) (*pinot.BrokerResponse, error) {
	table := "sales"
	query := "select storeCode, store, divisionCode, division, sum(totalPrice) as totalPrice from sales group by storeCode, divisionCode, store, division limit 1000000"
	brokerResp, err := r.db.ExecuteSQL(table, query)
	if err != nil {
		log.Error(err)
		return &pinot.BrokerResponse{}, err
	}

	return brokerResp, nil
}

func (r *salesRepository) GetSalesSummaryByStoreAndDivision(ctx context.Context, storeCode uint32) (*pinot.BrokerResponse, error) {
	table := "sales"
	query := fmt.Sprintf("select storeCode, store, divisionCode, division, sum(totalPrice) as totalPrice from sales where storeCode = %d group by storeCode, divisionCode, store, division limit 1000000", storeCode)
	brokerResp, err := r.db.ExecuteSQL(table, query)
	if err != nil {
		log.Error(err)
		return &pinot.BrokerResponse{}, err
	}

	return brokerResp, nil
}
