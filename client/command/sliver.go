package command

/*
	Sliver Implant Framework
	Copyright (C) 2023  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"

	"github.com/cs23m001/cipherc2/client/assets"
	"github.com/cs23m001/cipherc2/client/command/alias"
	"github.com/cs23m001/cipherc2/client/command/backdoor"
	"github.com/cs23m001/cipherc2/client/command/cursed"
	"github.com/cs23m001/cipherc2/client/command/dllhijack"
	"github.com/cs23m001/cipherc2/client/command/environment"
	"github.com/cs23m001/cipherc2/client/command/exec"
	"github.com/cs23m001/cipherc2/client/command/extensions"
	"github.com/cs23m001/cipherc2/client/command/filesystem"
	"github.com/cs23m001/cipherc2/client/command/info"
	"github.com/cs23m001/cipherc2/client/command/kill"
	"github.com/cs23m001/cipherc2/client/command/network"
	"github.com/cs23m001/cipherc2/client/command/pivots"
	"github.com/cs23m001/cipherc2/client/command/portfwd"
	"github.com/cs23m001/cipherc2/client/command/privilege"
	"github.com/cs23m001/cipherc2/client/command/processes"
	"github.com/cs23m001/cipherc2/client/command/reconfig"
	"github.com/cs23m001/cipherc2/client/command/registry"
	"github.com/cs23m001/cipherc2/client/command/rportfwd"
	"github.com/cs23m001/cipherc2/client/command/screenshot"
	"github.com/cs23m001/cipherc2/client/command/sessions"
	"github.com/cs23m001/cipherc2/client/command/shell"
	"github.com/cs23m001/cipherc2/client/command/socks"
	"github.com/cs23m001/cipherc2/client/command/tasks"
	"github.com/cs23m001/cipherc2/client/command/wasm"
	"github.com/cs23m001/cipherc2/client/command/wireguard"
	client "github.com/cs23m001/cipherc2/client/console"
	consts "github.com/cs23m001/cipherc2/client/constants"
	"github.com/reeflective/console"
	"github.com/spf13/cobra"
)

// SliverCommands returns all commands bound to the implant menu.
func SliverCommands(con *client.SliverClient) console.Commands {
	sliverCommands := func() *cobra.Command {
		sliver := &cobra.Command{
			Short: "Implant commands",
			CompletionOptions: cobra.CompletionOptions{
				HiddenDefaultCmd: true,
			},
		}

		// Utility function to be used for binding new commands to
		// the sliver menu: call the function with the name of the
		// group under which this/these commands should be added,
		// and the group will be automatically created if needed.
		bind := makeBind(sliver, con)

		// [ Core ]
		bind(consts.SliverCoreHelpGroup,
			reconfig.Commands,
			// sessions.Commands,
			sessions.SliverCommands,
			kill.Commands,
			// use.Commands,
			tasks.Commands,
			pivots.Commands,
		)

		// [ Info ]
		bind(consts.InfoHelpGroup,
			// info.Commands,
			info.SliverCommands,
			screenshot.Commands,
			environment.Commands,
			registry.Commands,
		)

		// [ Filesystem ]
		bind(consts.FilesystemHelpGroup,
			filesystem.Commands,
		)

		// [ Network tools ]
		bind(consts.NetworkHelpGroup,
			network.Commands,
			rportfwd.Commands,
			portfwd.Commands,
			socks.Commands,
			wireguard.SliverCommands,
		)

		// [ Execution ]
		bind(consts.ExecutionHelpGroup,
			shell.Commands,
			exec.Commands,
			backdoor.Commands,
			dllhijack.Commands,
			cursed.Commands,
			wasm.Commands,
		)

		// [ Privileges ]
		bind(consts.PrivilegesHelpGroup,
			privilege.Commands,
		)

		// [ Processes ]
		bind(consts.ProcessHelpGroup,
			processes.Commands,
		)

		// [ Aliases ]
		bind(consts.AliasHelpGroup)

		// [ Extensions ]
		bind(consts.ExtensionHelpGroup,
			extensions.Commands,
		)

		// [ Post-command declaration setup ]----------------------------------------

		// Load Aliases
		aliasManifests := assets.GetInstalledAliasManifests()
		for _, manifest := range aliasManifests {
			_, err := alias.LoadAlias(manifest, sliver, con)
			if err != nil {
				con.PrintErrorf("Failed to load alias: %s", err)
				continue
			}
		}

		// Load Extensions
		extensionManifests := assets.GetInstalledExtensionManifests()
		for _, manifest := range extensionManifests {
			mext, err := extensions.LoadExtensionManifest(manifest)
			// Absorb error in case there's no extensions manifest
			if err != nil {
				//con doesn't appear to be initialised here?
				//con.PrintErrorf("Failed to load extension: %s", err)
				fmt.Printf("Failed to load extension: %s\n", err)
				continue
			}

			for _, ext := range mext.ExtCommand {
				extensions.ExtensionRegisterCommand(ext, sliver, con)
			}
		}

		// [ Post-command declaration setup ]----------------------------------------

		// Everything below this line should preferably not be any command binding
		// (although you can do so without fear). If there are any final modifications
		// to make to the server menu command tree, it time to do them here.

		sliver.InitDefaultHelpCmd()
		sliver.SetHelpCommandGroupID(consts.SliverCoreHelpGroup)

		// Compute which commands should be available based on the current session/beacon.
		con.ExposeCommands()

		return sliver
	}

	return sliverCommands
}
