package filesystem

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2019  Bishop Fox

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
	"context"

	"github.com/cs23m001/CIPHERC2/client/console"
	"github.com/cs23m001/CIPHERC2/protobuf/CIPHERC2pb"
	"github.com/cs23m001/CIPHERC2/protobuf/clientpb"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// MkdirCmd - Make a remote directory.
func MkdirCmd(cmd *cobra.Command, con *console.CIPHERC2Client, args []string) {
	session, beacon := con.ActiveTarget.GetInteractive()
	if session == nil && beacon == nil {
		return
	}

	filePath := args[0]
	// filePath := ctx.Args.String("path")

	if filePath == "" {
		con.PrintErrorf("Missing parameter: directory name\n")
		return
	}

	mkdir, err := con.Rpc.Mkdir(context.Background(), &CIPHERC2pb.MkdirReq{
		Request: con.ActiveTarget.Request(cmd),
		Path:    filePath,
	})
	if err != nil {
		con.PrintErrorf("%s\n", err)
		return
	}
	if mkdir.Response != nil && mkdir.Response.Async {
		con.AddBeaconCallback(mkdir.Response.TaskID, func(task *clientpb.BeaconTask) {
			err = proto.Unmarshal(task.Response, mkdir)
			if err != nil {
				con.PrintErrorf("Failed to decode response %s\n", err)
				return
			}
			PrintMkdir(mkdir, con)
		})
		con.PrintAsyncResponse(mkdir.Response)
	} else {
		PrintMkdir(mkdir, con)
	}
}

// PrintMkdir - Print make directory.
func PrintMkdir(mkdir *CIPHERC2pb.Mkdir, con *console.CIPHERC2Client) {
	if mkdir.Response != nil && mkdir.Response.Err != "" {
		con.PrintErrorf("%s\n", mkdir.Response.Err)
		return
	}
	con.PrintInfof("%s\n", mkdir.Path)
}
