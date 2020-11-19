package handler

import (
	"encoding/json"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type AccountSummaryFlow struct {
	svc          accountSummaryService
	subAccountID string
}

func newGetAccountSummaryFlow(svc accountSummaryService, subAccountID string) *AccountSummaryFlow {
	return &AccountSummaryFlow{
		svc:          svc,
		subAccountID: subAccountID,
	}
}

func (f AccountSummaryFlow) doAccountSummaryRequest() goanda.AccountSummary {
	endpoint := "/accounts/" + f.subAccountID + "/summary"
	response := f.svc.OandaRequest(endpoint)
	res := goanda.AccountSummary{}
	json.Unmarshal(response, &res)
	return res
}
