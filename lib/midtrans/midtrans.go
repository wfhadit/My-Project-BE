package midtrans

import (
	"my-project-be/config"

	"github.com/veritrans/go-midtrans"
)

func GetMidtransClient(cfg *config.AppConfig) midtrans.Client {
	midtransClient := midtrans.NewClient()
	midtransClient.ServerKey = cfg.MIDTRANS_SERVER_KEY
	midtransClient.ClientKey = cfg.MIDTRANS_CLIENT_KEY
	midtransClient.APIEnvType = midtrans.Sandbox
	return midtransClient
}