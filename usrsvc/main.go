package main

import (
	"dogapm"
	"protos"
	"usrsvc/grpc"
)

func main() {
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/usrsvc"),
		dogapm.InfraRdsOption("127.0.0.1:6379"),
	)
	grpcserver := dogapm.NewGrpcServer(":8082")
	protos.RegisterUserServiceServer(grpcserver, &grpc.UserServer{})
	dogapm.EndPoint.Start()
}
