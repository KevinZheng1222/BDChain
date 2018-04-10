package client_test

import (
	"os"
	"testing"

	"github.com/bdc/abci/example/kvstore"

	nm "github.com/bdc/bdc/node"
	rpctest "github.com/bdc/bdc/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a bdc node (and merkleeyes) in the background to test against
	app := kvstore.NewKVStoreApplication()
	node = rpctest.Startbdc(app)
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}
