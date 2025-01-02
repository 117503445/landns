package rpclogic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/store"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LandnsServer struct {
	rpcgen.LanDNS
	store *store.LeasesStore
}

func (s *LandnsServer) SetLeases(ctx context.Context, req *rpcgen.SetLeasesRequest) (*emptypb.Empty, error) {
	s.store.SetLeases(req.AgentName, req.Leases)
	return &emptypb.Empty{}, nil
}

func NewLandnsServer(store *store.LeasesStore) *LandnsServer {
	return &LandnsServer{
		store: store,
	}
}

func (s *LandnsServer) Start(port int) {
	log.Info().Int("port", port).Msg("Starting landns server")
	twirpHandler := rpcgen.NewLanDNSServer(s)
	http.ListenAndServe(":"+fmt.Sprint(port), twirpHandler)
}
