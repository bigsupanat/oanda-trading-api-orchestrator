package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/conf"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/oandaclient"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/server"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
)

func main() {

	client, err := oandaclient.NewConnection(conf.AccountID, conf.AccessToken, conf.Environment)
	if err != nil {
		log.Fatalln(err)
	}

	svc := service.Service{
		OandaGetFunc:  client.Get,
		OandaPostFunc: client.Post,
		OandaPutFunc:  client.Put,
		//OandaRequestFunc: oanda.Request,
		// OandaSendFunc:              oanda.Send,
		// OandaGetOrderDetailsFunc:   oanda.GetOrderDetails,
		// OandaGetAccountSummaryFunc: oanda.GetAccountSummary,
		// OandaCreateOrderFunc:       oanda.CreateOrder,
	}

	serv := server.NewServer(svc)
	log.Println("Start server")
	serv.Start()

	defer func() {
		serv.Stop()
		log.Println("Stop server successfully")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")

}
