package conf

import "os"

var (
	ServicePort = os.Getenv("SERVICE_PORT")
	Environment = os.Getenv("ENVIRONMENT")
	AccountID   = os.Getenv("ACCOUNT_ID")
	AccessToken = os.Getenv("ACCESS_TOKEN")
)
