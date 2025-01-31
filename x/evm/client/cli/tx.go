package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/palomachain/paloma/v2/x/evm/types"
	"github.com/spf13/cobra"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRemoveSmartContractDeployment())
	cmd.AddCommand(CmdUploadUserSmartContract())
	cmd.AddCommand(CmdRemoveUserSmartContract())
	cmd.AddCommand(CmdDeployUserSmartContract())

	return cmd
}
