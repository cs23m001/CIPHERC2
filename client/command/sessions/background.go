package sessions

import (
	"github.com/cs23m001/cipherc2/client/console"
	"github.com/spf13/cobra"
)

// BackgroundCmd - Background the active session.
func BackgroundCmd(cmd *cobra.Command, con *console.SliverClient, args []string) {
	con.ActiveTarget.Background()
	con.PrintInfof("Background ...\n")
}
