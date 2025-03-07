package rpc

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2021  Bishop Fox

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
)

// RegistryRead - gRPC interface to read a registry key from a session
func (rpc *Server) RegistryRead(ctx context.Context, req *CIPHERC2pb.RegistryReadReq) (*CIPHERC2pb.RegistryRead, error) {
	resp := &CIPHERC2pb.RegistryRead{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryWrite - gRPC interface to write to a registry key on a session
func (rpc *Server) RegistryWrite(ctx context.Context, req *CIPHERC2pb.RegistryWriteReq) (*CIPHERC2pb.RegistryWrite, error) {
	resp := &CIPHERC2pb.RegistryWrite{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryCreateKey - gRPC interface to create a registry key on a session
func (rpc *Server) RegistryCreateKey(ctx context.Context, req *CIPHERC2pb.RegistryCreateKeyReq) (*CIPHERC2pb.RegistryCreateKey, error) {
	resp := &CIPHERC2pb.RegistryCreateKey{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryDeleteKey - gRPC interface to delete a registry key on a session
func (rpc *Server) RegistryDeleteKey(ctx context.Context, req *CIPHERC2pb.RegistryDeleteKeyReq) (*CIPHERC2pb.RegistryDeleteKey, error) {
	resp := &CIPHERC2pb.RegistryDeleteKey{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryListSubKeys - gRPC interface to list the sub keys of a registry key
func (rpc *Server) RegistryListSubKeys(ctx context.Context, req *CIPHERC2pb.RegistrySubKeyListReq) (*CIPHERC2pb.RegistrySubKeyList, error) {
	resp := &CIPHERC2pb.RegistrySubKeyList{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryListSubKeys - gRPC interface to list the sub keys of a registry key
func (rpc *Server) RegistryListValues(ctx context.Context, req *CIPHERC2pb.RegistryListValuesReq) (*CIPHERC2pb.RegistryValuesList, error) {
	resp := &CIPHERC2pb.RegistryValuesList{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegistryDumpHive - gRPC interface to dump a specific registry hive as a binary file
func (rpc *Server) RegistryReadHive(ctx context.Context, req *CIPHERC2pb.RegistryReadHiveReq) (*CIPHERC2pb.RegistryReadHive, error) {
	resp := &CIPHERC2pb.RegistryReadHive{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
