package server

import (
	"encoding/json"
	"net/http"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/handler"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
)

// func getAccountSummary(w http.ResponseWriter, r *http.Request) {
// 	// res, err := json.Marshal
// }

func getAccountSummaryFn(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := "001-011-2498245-001"
		res := handler.GetAccountSummary(s, req)
		bytRet, _ := json.Marshal(res)

		w.WriteHeader(http.StatusOK)
		w.Write(bytRet)

	}
}
