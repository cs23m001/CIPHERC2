package licenses

import (
	"github.com/spf13/cobra"

	"github.com/cs23m001/cipherc2/client/command/help"
	"github.com/cs23m001/cipherc2/client/console"
	consts "github.com/cs23m001/cipherc2/client/constants"
	"github.com/cs23m001/cipherc2/client/licenses"
)

// Commands returns the `licences` command.
func Commands(con *console.SliverClient) []*cobra.Command {
	licensesCmd := &cobra.Command{
		Use:   consts.LicensesStr,
		Short: "Open source licenses",
		Long:  help.GetHelpFor([]string{consts.LicensesStr}),
		Run: func(cmd *cobra.Command, args []string) {
			con.Println(licenses.All)
		},
		GroupID: consts.GenericHelpGroup,
	}

	return []*cobra.Command{licensesCmd}
}
