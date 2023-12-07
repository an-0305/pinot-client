package util

import (
	"encoding/json"

	"github.com/startreedata/pinot-client-go/pinot"
)

func GetStringValue(table *pinot.ResultTable, row []interface{}, columnName string) string {
	for i, name := range table.DataSchema.ColumnNames {
		if name == columnName {
			return row[i].(string)
		}
	}
	return ""
}

func GetInt32Value(table *pinot.ResultTable, row []interface{}, columnName string) int32 {
	for i, name := range table.DataSchema.ColumnNames {
		if name == columnName {
			val, _ := row[i].(json.Number)
			float64Val, err := val.Float64()
			if err != nil {
				return 0
			}
			return int32(float64Val)
		}
	}
	return 0
}

func CreateJSONResponse(brokerResp *pinot.BrokerResponse, mappingFunc func(*pinot.ResultTable, []interface{}) interface{}) []interface{} {
	var outputDataSlice []interface{}
	for _, row := range brokerResp.ResultTable.Rows {
		data := mappingFunc(brokerResp.ResultTable, row)
		outputDataSlice = append(outputDataSlice, data)
	}
	return outputDataSlice
}
