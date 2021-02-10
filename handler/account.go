package handler

import (
	_ "context"
	"encoding/json"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type accountService interface {
	Request(endpoint string) []byte
}

func GetAccountSummary(svc accountService, accountID string) (goanda.AccountSummary, error) {
	endpoint := "/accounts/" + accountID + "/summary"
	response := svc.Request(endpoint)
	data := goanda.AccountSummary{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
