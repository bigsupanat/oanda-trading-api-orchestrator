package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/goanda"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/server"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
	"github.com/spf13/viper"
)

type OandaAPIType struct {
	Key          string `mapstructure:"key"`
	AccountID    string `mapstructure:"accountid"`
	Live         bool   `mapstructure:"live"`
	SubAccountID string `mapstructure:"subaccountid"`
}

func main() {

	var oandaAPI OandaAPIType

	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&oandaAPI)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	oanda := goanda.NewConnection(oandaAPI.AccountID, oandaAPI.Key, oandaAPI.Live)

	svc := service.Service{
		OandaRequestFunc:           oanda.Request,
		OandaSendFunc:              oanda.Send,
		OandaUpdateFunc:            oanda.Update,
		OandaGetOrderDetailsFunc:   oanda.GetOrderDetails,
		OandaGetAccountSummaryFunc: oanda.GetAccountSummary,
		OandaCreateOrderFunc:       oanda.CreateOrder,
	}

	serv := server.NewServer(svc)

	log.Println("Start server")
	serv.Start()

	defer func() {
		serv.Stop()
		log.Println("Stop server successfully")

	}()

	// accountSummary := oanda.GetAccountSummary(oandaAPI.SubAccountID)
	// spew.Dump(accountSummary)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")

}
