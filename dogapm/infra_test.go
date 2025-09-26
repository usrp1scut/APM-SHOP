package dogapm

import (
	"context"
	"fmt"
	"net/http"
	"protos"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Infra.Init(InfraDbOption("root:root@tcp(127.0.0.1:3306)/ordersvc"),
		InfraRdsOption("127.0.0.1:6379"))

}

func TestNewHttpServer(t *testing.T) {
	s := NewHttpServer(":8080")
	s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	s.Start()
	time.Sleep(time.Hour)
}

type helloSvc struct {
	protos.UnimplementedHelloServiceServer
}

func (h *helloSvc) Receive(ctx context.Context, msg *protos.HelloMsg) (*protos.HelloMsg, error) {
	return msg, nil
}

func TestGrpc(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8080")
		protos.RegisterHelloServiceServer(s, &helloSvc{})
		s.Start()
	}()
	client := NewGrpcClient("127.0.0.1:8080")
	res, err := protos.NewHelloServiceClient(client).Receive(context.TODO(), &protos.HelloMsg{Msg: "hello world"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Msg)
}
