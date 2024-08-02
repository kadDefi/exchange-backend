package main

import (
	"context"
	"exchange-backend/cmd/root"
	"exchange-backend/internal/pkg/log"
	"github.com/spf13/viper"
)

var Version string
var cfgFile string

func main() {
	logger := log.FromContext(context.Background())

	if err := root.Cmd.Execute(); err != nil {
		logger.Sugar().Fatal(err)
	}
}

func init() {
	root.Cmd.PersistentFlags().String("config", "./configs/local.yml", "config file (default is ./configs/local.yml)")
	viper.BindPFlag("config", root.Cmd.PersistentFlags().Lookup("config"))
}
