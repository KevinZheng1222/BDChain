package main

import (
	"fmt"
	"net/http"
	"os"

	rpcserver "github.com/bdc/bdc/rpc/lib/server"
	cmn "github.com/bdc/tmlibs/common"
	"github.com/bdc/tmlibs/log"
)

var routes = map[string]*rpcserver.RPCFunc{
	"hello_world": rpcserver.NewRPCFunc(HelloWorld, "name,num"),
}

func HelloWorld(name string, num int) (Result, error) {
	return Result{fmt.Sprintf("hi %s %d", name, num)}, nil
}

type Result struct {
	Result string
}

func main() {
	mux := http.NewServeMux()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	rpcserver.RegisterRPCFuncs(mux, routes, logger)
	_, err := rpcserver.StartHTTPServer("0.0.0.0:8008", mux, logger)
	if err != nil {
		cmn.Exit(err.Error())
	}

	// Wait forever
	cmn.TrapSignal(func() {
	})

}
