package rpc

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

	"github.com/cs23m001/CIPHERC2/protobuf/CIPHERC2pb"
	"github.com/cs23m001/CIPHERC2/protobuf/commonpb"
	"github.com/cs23m001/CIPHERC2/server/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	// ErrTunnelInitFailure - Returned when a tunnel cannot be initialized
	ErrTunnelInitFailure = status.Error(codes.Internal, "Failed to initialize tunnel")
)

// Shell - Open an interactive shell
func (rpc *Server) Shell(ctx context.Context, req *CIPHERC2pb.ShellReq) (*CIPHERC2pb.Shell, error) {
	session := core.Sessions.Get(req.Request.SessionID)
	if session == nil {
		return nil, ErrInvalidSessionID
	}
	tunnel := core.Tunnels.Get(req.TunnelID)
	if tunnel == nil {
		return nil, core.ErrInvalidTunnelID
	}
	reqData, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	data, err := session.Request(CIPHERC2pb.MsgNumber(req), rpc.getTimeout(req), reqData)
	if err != nil {
		return nil, err
	}
	shell := &CIPHERC2pb.Shell{}
	err = proto.Unmarshal(data, shell)
	return shell, err
}

// RunSSHCommand runs a SSH command using the client built into the implant
func (rpc *Server) RunSSHCommand(ctx context.Context, req *CIPHERC2pb.SSHCommandReq) (*CIPHERC2pb.SSHCommand, error) {
	resp := &CIPHERC2pb.SSHCommand{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
