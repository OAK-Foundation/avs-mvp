package cmd

import (
	"github.com/spf13/cobra"

	"github.com/OAK-Foundation/oak-avs/operator"
)

var (
	deRegisterCmd = &cobra.Command{
		Use:   "deregister",
		Short: "DeRegister your operator from Oak AVS",
		Long:  `Opt-out your operator from AVS.\nThe process will failed if you haven't opt-in yet`,
		Run: func(cmd *cobra.Command, args []string) {
			operator.DeregisterFromAVS(config)
		},
	}
)

func init() {
	rootCmd.AddCommand(deRegisterCmd)
}
