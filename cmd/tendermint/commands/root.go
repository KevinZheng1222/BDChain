package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/bdc/bdc/config"
	"github.com/bdc/tmlibs/cli"
	tmflags "github.com/bdc/tmlibs/cli/flags"
	"github.com/bdc/tmlibs/log"
)

var (
	config = cfg.DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

func init() {
	registerFlagsRootCmd(RootCmd)
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("log_level", config.LogLevel, "Log level")
}

// ParseConfig retrieves the default environment configuration,
// sets up the bdc root and ensures that the root exists
func ParseConfig() (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	conf.SetRoot(conf.RootDir)
	cfg.EnsureRoot(conf.RootDir)
	return conf, err
}

// RootCmd is the root command for bdc core.
var RootCmd = &cobra.Command{
	Use:   "bdc",
	Short: "bdc Core (BFT Consensus) in Go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}
		config, err = ParseConfig()
		if err != nil {
			return err
		}
		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
		if err != nil {
			return err
		}
		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}
		logger = logger.With("module", "main")
		return nil
	},
}
