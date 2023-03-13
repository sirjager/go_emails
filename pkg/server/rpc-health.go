package server

import (
	"context"
	"time"

	rpc "github.com/sirjager/rpcs/emails/go"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (srvr *RPCEmailsServer) EmailsHealth(ctx context.Context, req *rpc.EmailsHealthRequest) (*rpc.EmailsHealthResponse, error) {
	res := &rpc.EmailsHealthResponse{
		Status:    "UP",
		Timestamp: timestamppb.New(time.Now()),
		Protected: srvr.config.ApiSecret != "",
		Uptime:    durationpb.New(time.Since(srvr.startTime)),
		Started:   timestamppb.New(srvr.startTime),
	}
	return res, nil
}
