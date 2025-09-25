package dogapm

import "testing"

func TestInit(t *testing.T) {
	Infra.Init(InfraDbOption("root:root@tcp(127.0.0.1:3306)/ordersvc"),
		InfraRdsOption("127.0.0.1:6379"))

}
