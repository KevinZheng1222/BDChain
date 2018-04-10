package rpctest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bdc/tmlibs/log"

	abci "github.com/bdc/abci/types"
	cmn "github.com/bdc/tmlibs/common"

	cfg "github.com/bdc/bdc/config"
	nm "github.com/bdc/bdc/node"
	"github.com/bdc/bdc/proxy"
	ctypes "github.com/bdc/bdc/rpc/core/types"
	core_grpc "github.com/bdc/bdc/rpc/grpc"
	rpcclient "github.com/bdc/bdc/rpc/lib/client"
	"github.com/bdc/bdc/types"
)

var globalConfig *cfg.Config

func waitForRPC() {
	laddr := GetConfig().RPC.ListenAddress
	client := rpcclient.NewJSONRPCClient(laddr)
	result := new(ctypes.ResultStatus)
	for {
		_, err := client.Call("status", map[string]interface{}{}, result)
		if err == nil {
			return
		}
	}
}

func waitForGRPC() {
	client := GetGRPCClient()
	for {
		_, err := client.Ping(context.Background(), &core_grpc.RequestPing{})
		if err == nil {
			return
		}
	}
}

// f**ing long, but unique for each test
func makePathname() string {
	// get path
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// fmt.Println(p)
	sep := string(filepath.Separator)
	return strings.Replace(p, sep, "_", -1)
}

func randPort() int {
	return int(cmn.RandUint16()/2 + 10000)
}

func makeAddrs() (string, string, string) {
	start := randPort()
	return fmt.Sprintf("tcp://0.0.0.0:%d", start),
		fmt.Sprintf("tcp://0.0.0.0:%d", start+1),
		fmt.Sprintf("tcp://0.0.0.0:%d", start+2)
}

// GetConfig returns a config for the test cases as a singleton
func GetConfig() *cfg.Config {
	if globalConfig == nil {
		pathname := makePathname()
		globalConfig = cfg.ResetTestRoot(pathname)

		// and we use random ports to run in parallel
		tm, rpc, grpc := makeAddrs()
		globalConfig.P2P.ListenAddress = tm
		globalConfig.RPC.ListenAddress = rpc
		globalConfig.RPC.GRPCListenAddress = grpc
		globalConfig.TxIndex.IndexTags = "app.creator" // see kvstore application
	}
	return globalConfig
}

func GetGRPCClient() core_grpc.BroadcastAPIClient {
	grpcAddr := globalConfig.RPC.GRPCListenAddress
	return core_grpc.StartGRPCClient(grpcAddr)
}

// Startbdc starts a test bdc server in a go routine and returns when it is initialized
func Startbdc(app abci.Application) *nm.Node {
	node := Newbdc(app)
	err := node.Start()
	if err != nil {
		panic(err)
	}

	// wait for rpc
	waitForRPC()
	waitForGRPC()

	fmt.Println("bdc running!")

	return node
}

// Newbdc creates a new bdc server and sleeps forever
func Newbdc(app abci.Application) *nm.Node {
	// Create & start node
	config := GetConfig()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewFilter(logger, log.AllowError())
	privValidatorFile := config.PrivValidatorFile()
	privValidator := types.LoadOrGenPrivValidatorFS(privValidatorFile)
	papp := proxy.NewLocalClientCreator(app)
	node, err := nm.NewNode(config, privValidator, papp,
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider, logger)
	if err != nil {
		panic(err)
	}
	return node
}
