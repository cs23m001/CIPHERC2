package rpc

import (
	"context"

	"github.com/cs23m001/CIPHERC2/protobuf/CIPHERC2pb"
	"github.com/cs23m001/CIPHERC2/protobuf/commonpb"
)

// Services - List and control services
func (rpc *Server) Services(ctx context.Context, req *CIPHERC2pb.ServicesReq) (*CIPHERC2pb.Services, error) {
	resp := &CIPHERC2pb.Services{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rpc *Server) ServiceDetail(ctx context.Context, req *CIPHERC2pb.ServiceDetailReq) (*CIPHERC2pb.ServiceDetail, error) {
	resp := &CIPHERC2pb.ServiceDetail{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StartService creates and starts a Windows service on a remote host
func (rpc *Server) StartService(ctx context.Context, req *CIPHERC2pb.StartServiceReq) (*CIPHERC2pb.ServiceInfo, error) {
	resp := &CIPHERC2pb.ServiceInfo{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rpc *Server) StartServiceByName(ctx context.Context, req *CIPHERC2pb.StartServiceByNameReq) (*CIPHERC2pb.ServiceInfo, error) {
	resp := &CIPHERC2pb.ServiceInfo{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StopService stops a remote service
func (rpc *Server) StopService(ctx context.Context, req *CIPHERC2pb.StopServiceReq) (*CIPHERC2pb.ServiceInfo, error) {
	resp := &CIPHERC2pb.ServiceInfo{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RemoveService deletes a service from the remote system
func (rpc *Server) RemoveService(ctx context.Context, req *CIPHERC2pb.RemoveServiceReq) (*CIPHERC2pb.ServiceInfo, error) {
	resp := &CIPHERC2pb.ServiceInfo{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
