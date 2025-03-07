package sessions

import (
	"github.com/cs23m001/CIPHERC2/client/console"
	"github.com/spf13/cobra"
)

// BackgroundCmd - Background the active session.
func BackgroundCmd(cmd *cobra.Command, con *console.CIPHERC2Client, args []string) {
	con.ActiveTarget.Background()
	con.PrintInfof("Background ...\n")
}
