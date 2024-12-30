package rpclogic

import (
	"context"

	"github.com/117503445/landns/pkg/rpcgen"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LandnsServer struct {
	rpcgen.LanDNS
}

// SetLeases(context.Context, *SetLeasesRequest) (*google_protobuf.Empty, error)

func (s *LandnsServer) SetLeases(ctx context.Context, req *rpcgen.SetLeasesRequest) (*emptypb.Empty, error) {
	return nil, nil
}
