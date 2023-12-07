package db

import (
	"github.com/labstack/gommon/log"
	"github.com/startreedata/pinot-client-go/pinot"
)

func Init() (*pinot.Connection, error) {
	pinotClient, err := pinot.NewFromBrokerList([]string{"pinot-broker:8099"})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return pinotClient, nil
}
