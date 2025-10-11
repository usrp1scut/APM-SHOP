package dogapm

import (
	"os"
	"os/signal"
	"syscall"
)

type starter interface {
	Start()
}

type closer interface {
	Close()
}

var (
	globalStarters = make([]starter, 0)
	globalClosers  = make([]closer, 0)
)

type endPoint struct {
	stop chan int
}

var EndPoint = &endPoint{stop: make(chan int, 1)}

func (e *endPoint) Start() {
	for _, com := range globalStarters {
		com.Start()
	}
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		e.Shutdown()
	}()
	<-e.stop
}

func (e *endPoint) Shutdown() {
	for _, com := range globalClosers {
		com.Close()
	}
	e.stop <- 1
}
