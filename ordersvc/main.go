package main

import (
	"dogapm"
	"net/http"
	"ordersvc/api"
	"ordersvc/grpcclient"
	"protos"
)

func main() {
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/ordersvc"),
	)

	grpcclient.SkuClient = protos.NewSkuServiceClient(dogapm.NewGrpcClient(":8081"))
	grpcclient.UserClient = protos.NewUserServiceClient(dogapm.NewGrpcClient(":8082"))

	httpserver := dogapm.NewHttpServer(":8080")
	httpserver.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	httpserver.HandleFunc("/order/add", api.Order.Add)
	dogapm.EndPoint.Start()
}
