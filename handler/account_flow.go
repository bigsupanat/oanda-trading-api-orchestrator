package handler

import (
	"encoding/json"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
)

type AccountSummaryFlow struct {
	svc          accountService
	subAccountID string
}

func newGetAccountSummaryFlow(svc accountService, subAccountID string) *AccountSummaryFlow {
	return &AccountSummaryFlow{
		svc:          svc,
		subAccountID: subAccountID,
	}
}

func (f AccountSummaryFlow) doAccountSummaryRequest() goanda.AccountSummary {
	endpoint := "/accounts/" + f.subAccountID + "/summary"
	response := f.svc.Request(endpoint)
	res := goanda.AccountSummary{}
	json.Unmarshal(response, &res)
	return res
}
