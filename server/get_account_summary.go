package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/conf"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/handler"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
)

func getAccountSummaryFn(s service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := handler.GetAccountSummary(s, conf.AccountID)
		if err != nil {
			log.Println(err)
		}
		bytRet, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytRet)
	}
}
