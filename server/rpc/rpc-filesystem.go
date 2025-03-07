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
	"crypto/sha256"
	"fmt"

	"github.com/cs23m001/CIPHERC2/protobuf/CIPHERC2pb"
	"github.com/cs23m001/CIPHERC2/protobuf/commonpb"
	"github.com/cs23m001/CIPHERC2/server/core"
	"github.com/cs23m001/CIPHERC2/server/db"
	"github.com/cs23m001/CIPHERC2/server/db/models"
	"github.com/cs23m001/CIPHERC2/server/log"
	"github.com/cs23m001/CIPHERC2/util/encoders"
)

var (
	fsLog = log.NamedLogger("rcp", "fs")
)

// Ls - List a directory
func (rpc *Server) Ls(ctx context.Context, req *CIPHERC2pb.LsReq) (*CIPHERC2pb.Ls, error) {
	resp := &CIPHERC2pb.Ls{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Mv - Move or rename a file
func (rpc *Server) Mv(ctx context.Context, req *CIPHERC2pb.MvReq) (*CIPHERC2pb.Mv, error) {
	resp := &CIPHERC2pb.Mv{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Cp - Copy a file to another location
func (rpc *Server) Cp(ctx context.Context, req *CIPHERC2pb.CpReq) (*CIPHERC2pb.Cp, error) {
	resp := &CIPHERC2pb.Cp{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Rm - Remove file or directory
func (rpc *Server) Rm(ctx context.Context, req *CIPHERC2pb.RmReq) (*CIPHERC2pb.Rm, error) {
	resp := &CIPHERC2pb.Rm{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Mkdir - Make a directory
func (rpc *Server) Mkdir(ctx context.Context, req *CIPHERC2pb.MkdirReq) (*CIPHERC2pb.Mkdir, error) {
	resp := &CIPHERC2pb.Mkdir{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Cd - Change directory
func (rpc *Server) Cd(ctx context.Context, req *CIPHERC2pb.CdReq) (*CIPHERC2pb.Pwd, error) {
	resp := &CIPHERC2pb.Pwd{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Pwd - Print working directory
func (rpc *Server) Pwd(ctx context.Context, req *CIPHERC2pb.PwdReq) (*CIPHERC2pb.Pwd, error) {
	resp := &CIPHERC2pb.Pwd{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Download - Download a file from the remote file system
func (rpc *Server) Download(ctx context.Context, req *CIPHERC2pb.DownloadReq) (*CIPHERC2pb.Download, error) {
	resp := &CIPHERC2pb.Download{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Upload - Upload a file from the remote file system
func (rpc *Server) Upload(ctx context.Context, req *CIPHERC2pb.UploadReq) (*CIPHERC2pb.Upload, error) {
	resp := &CIPHERC2pb.Upload{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	if req.IsIOC {
		go trackIOC(req, resp)
	}
	return resp, nil
}

// Chmod - Change permission on a file or directory
func (rpc *Server) Chmod(ctx context.Context, req *CIPHERC2pb.ChmodReq) (*CIPHERC2pb.Chmod, error) {
	resp := &CIPHERC2pb.Chmod{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Chown - Change owner on a file or directory
func (rpc *Server) Chown(ctx context.Context, req *CIPHERC2pb.ChownReq) (*CIPHERC2pb.Chown, error) {
	resp := &CIPHERC2pb.Chown{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Chtimes - Change file access and modification times on a file or directory
func (rpc *Server) Chtimes(ctx context.Context, req *CIPHERC2pb.ChtimesReq) (*CIPHERC2pb.Chtimes, error) {
	resp := &CIPHERC2pb.Chtimes{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemfilesList - List memfiles
func (rpc *Server) MemfilesList(ctx context.Context, req *CIPHERC2pb.MemfilesListReq) (*CIPHERC2pb.Ls, error) {
	resp := &CIPHERC2pb.Ls{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemfilesAdd - Add memfile
func (rpc *Server) MemfilesAdd(ctx context.Context, req *CIPHERC2pb.MemfilesAddReq) (*CIPHERC2pb.MemfilesAdd, error) {
	resp := &CIPHERC2pb.MemfilesAdd{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemfilesRm - Close memfile
func (rpc *Server) MemfilesRm(ctx context.Context, req *CIPHERC2pb.MemfilesRmReq) (*CIPHERC2pb.MemfilesRm, error) {
	resp := &CIPHERC2pb.MemfilesRm{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func hashUploadData(encoder string, data []byte) [32]byte {
	if encoder == "gzip" {
		decodedData, err := new(encoders.Gzip).Decode(data)
		if err != nil {
			return sha256.Sum256(nil)
		}
		return sha256.Sum256(decodedData)
	} else {
		return sha256.Sum256(data)
	}
}

func trackIOC(req *CIPHERC2pb.UploadReq, resp *CIPHERC2pb.Upload) {
	fsLog.Debugf("Adding IOC to database ...")
	request := req.GetRequest()
	if request == nil {
		fsLog.Error("No request for upload")
		return
	}
	session := core.Sessions.Get(request.SessionID)
	if session == nil {
		fsLog.Error("No session for upload request")
		return
	}
	host, err := db.HostByHostUUID(session.UUID)
	if err != nil {
		fsLog.Errorf("No host for session uuid %v", session.UUID)
		return
	}

	sum := hashUploadData(req.Encoder, req.Data)
	ioc := &models.IOC{
		HostID:   host.HostUUID,
		Path:     resp.Path,
		FileHash: fmt.Sprintf("%x", sum),
	}
	if db.Session().Create(ioc).Error != nil {
		fsLog.Error("Failed to create IOC")
	}
}

// Grep - Search a file or directory for text matching a regex
func (rpc *Server) Grep(ctx context.Context, req *CIPHERC2pb.GrepReq) (*CIPHERC2pb.Grep, error) {
	resp := &CIPHERC2pb.Grep{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Mount - Get information on mounted filesystems
func (rpc *Server) Mount(ctx context.Context, req *CIPHERC2pb.MountReq) (*CIPHERC2pb.Mount, error) {
	resp := &CIPHERC2pb.Mount{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
