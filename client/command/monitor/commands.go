package monitor

import (
	"github.com/spf13/cobra"

	"github.com/cs23m001/CIPHERC2/client/console"
	consts "github.com/cs23m001/CIPHERC2/client/constants"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.CIPHERC2Client) []*cobra.Command {
	monitorCmd := &cobra.Command{
		Use:     consts.MonitorStr,
		Short:   "Monitor threat intel platforms for CIPHERC2 implants",
		GroupID: consts.CIPHERC2HelpGroup,
	}
	monitorCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the monitoring loops",
		Run: func(cmd *cobra.Command, args []string) {
			MonitorStartCmd(cmd, con, args)
		},
	})
	monitorCmd.AddCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stop the monitoring loops",
		Run: func(cmd *cobra.Command, args []string) {
			MonitorStopCmd(cmd, con, args)
		},
	})

	return []*cobra.Command{monitorCmd}
}
