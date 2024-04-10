package main

import (
	net_http "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/common-library/go/http"
	"github.com/common-library/go/log/klog"
)

func main() {
	defer klog.Flush()

	klog.InfoS("main start")
	defer klog.InfoS("main end")

	server := http.Server{}

	livez := func(w net_http.ResponseWriter, r *net_http.Request) { w.WriteHeader(net_http.StatusOK) }
	server.RegisterHandlerFunc("/livez", net_http.MethodGet, livez)

	readyz := func(w net_http.ResponseWriter, r *net_http.Request) { w.WriteHeader(net_http.StatusOK) }
	server.RegisterHandlerFunc("/readyz", net_http.MethodGet, readyz)

	listenAndServeFailureFunc := func(err error) { klog.ErrorS(err, "") }
	if err := server.Start(":10000", listenAndServeFailureFunc); err != nil {
		klog.ErrorS(err, "")
	}
	defer server.Stop(30)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	klog.InfoS("signal", "kind", <-signals)
}
