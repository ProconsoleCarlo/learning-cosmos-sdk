package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"hello/x/hello/types"
)

var _ = strconv.Itoa(0)

func CmdSayEvenOrOdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "say-even-or-odd [number]",
		Short: "Query say-even-or-odd",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqNumber := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySayEvenOrOddRequest{

				Number: reqNumber,
			}

			res, err := queryClient.SayEvenOrOdd(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
