package rpclogic

import (
	"context"

	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LandnsServer struct {
	rpcgen.LanDNS
	store *store.LeasesStore
}


func (s *LandnsServer) SetLeases(ctx context.Context, req *rpcgen.SetLeasesRequest) (*emptypb.Empty, error) {
	s.store.SetLeases(req.AgentName, req.Leases)
	return nil, nil
}

func NewLandnsServer(store *store.LeasesStore) *LandnsServer {
	return &LandnsServer{
		store: store,
	}
}