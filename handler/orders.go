package handler

import (
	"encoding/json"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/conf"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type ordersService interface {
	Send(endpoint string, data []byte) []byte
	Update(endpoint string, data []byte) []byte
}

func CreateOrder(svc ordersService, body goanda.OrderPayload) (goanda.OrderResponse, error) {
	endpoint := "/accounts/" + conf.AccountID + "/orders"
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return goanda.OrderResponse{}, err
	}
	response := svc.Send(endpoint, jsonBody)
	data := goanda.OrderResponse{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return goanda.OrderResponse{}, nil
	}
	return data, nil
}

func CancelOrder(svc ordersService, orderSpecifier string) (goanda.CancelledOrder, error) {
	endpoint := "/accounts/" + conf.AccountID + "/orders/" + orderSpecifier + "/cancel"
	response := svc.Update(endpoint, nil)
	data := goanda.CancelledOrder{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		return goanda.CancelledOrder{}, nil
	}
	return data, nil
}
