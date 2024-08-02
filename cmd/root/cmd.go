package root

import (
	"exchange-backend/cmd/script"
	"exchange-backend/cmd/unit"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "exchange-backend",
	Short: "exchange-backend is a  backend for NFT Exchange",
}

func init() {
	Cmd.AddCommand(unit.Cmd)
	Cmd.AddCommand(script.Cmd)
}
