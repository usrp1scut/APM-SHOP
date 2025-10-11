package main

import (
	"dogapm"
	"protos"
	"skusvc/grpc"
)

func main() {
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/skusvc"),
	)
	grpcserver := dogapm.NewGrpcServer(":8081")
	protos.RegisterSkuServiceServer(grpcserver, &grpc.SkuServer{})
	dogapm.EndPoint.Start()
}
