package rpc

import (
	"context"

	"github.com/cs23m001/cipherc2/protobuf/commonpb"
	"github.com/cs23m001/cipherc2/protobuf/cipherc2pb"
)

// Ping - Try to send a round trip message to the implant
func (rpc *Server) Ping(ctx context.Context, req *sliverpb.Ping) (*sliverpb.Ping, error) {
	resp := &sliverpb.Ping{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
