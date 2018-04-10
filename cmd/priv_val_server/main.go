package main

import (
	"flag"
	"os"

	crypto "github.com/bdc/go-crypto"
	cmn "github.com/bdc/tmlibs/common"
	"github.com/bdc/tmlibs/log"

	priv_val "github.com/bdc/bdc/types/priv_validator"
)

func main() {
	var (
		addr        = flag.String("addr", ":46659", "Address of client to connect to")
		chainID     = flag.String("chain-id", "mychain", "chain id")
		privValPath = flag.String("priv", "", "priv val file path")

		logger = log.NewTMLogger(
			log.NewSyncWriter(os.Stdout),
		).With("module", "priv_val")
	)
	flag.Parse()

	logger.Info(
		"Starting private validator",
		"addr", *addr,
		"chainID", *chainID,
		"privPath", *privValPath,
	)

	privVal := priv_val.LoadPrivValidatorJSON(*privValPath)

	rs := priv_val.NewRemoteSigner(
		logger,
		*chainID,
		*addr,
		privVal,
		crypto.GenPrivKeyEd25519(),
	)
	err := rs.Start()
	if err != nil {
		panic(err)
	}

	cmn.TrapSignal(func() {
		err := rs.Stop()
		if err != nil {
			panic(err)
		}
	})
}
