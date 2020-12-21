package grpc

import (
	"context"
	"net"

	emailProto "github.com/NODO-UH/uh-email-quota-protob"
	"github.com/NODO-UH/uh-email-quota/src/quota"
	"google.golang.org/grpc"
)

type EmailQuotaRPC struct {
	emailProto.EmalQuotaServer
}

func (eq *EmailQuotaRPC) GetQuota(ctx context.Context, ud *emailProto.UserData) (*emailProto.UserQuota, error) {
	if quotaInfo, err := quota.GetUserQuota(ud.User); err != nil {
		return nil, err
	} else {
		return &emailProto.UserQuota{
			Value: quotaInfo.Value,
			Limit: quotaInfo.Limit,
		}, nil
	}
}

func StartGRPC() error {
	if lis, err := net.Listen("tcp", "localhost:8000"); err != nil {
		return err
	} else {
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		emailProto.RegisterEmalQuotaServer(grpcServer, &EmailQuotaRPC{})
		return grpcServer.Serve(lis)
	}
}
