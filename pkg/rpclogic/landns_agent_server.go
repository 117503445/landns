package rpclogic

import "fmt"

import (
	"context"
	"net/http"
	"sync"

	"github.com/117503445/landns/pkg/rpcgen"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LandnsAgentServer struct {
	rpcgen.LanDNSAgent

	// leaseChan     <-chan []*rpcgen.Lease
	curLeases     []*rpcgen.Lease
	curLeasesLock sync.RWMutex
}

// GetLeases(context.Context, *google_protobuf.Empty) (*GetLeasesResponse, error)

func (s *LandnsAgentServer) GetLeases(context.Context, *emptypb.Empty) (*rpcgen.GetLeasesResponse, error) {
	return &rpcgen.GetLeasesResponse{
		Leases: s.curLeases,
	}, nil
}

func NewLandnsAgentServer() *LandnsAgentServer {
	return &LandnsAgentServer{
		// leaseChan: leaseChan,
		curLeases: make([]*rpcgen.Lease, 0),
	}
}

func (s *LandnsAgentServer) SetLeases(leases []*rpcgen.Lease) {
	s.curLeasesLock.Lock()
	defer s.curLeasesLock.Unlock()
	s.curLeases = leases
}

func (s *LandnsAgentServer) Start(port int) {
	twirpHandler := rpcgen.NewLanDNSAgentServer(s)
	http.ListenAndServe(":"+fmt.Sprint(port), twirpHandler)
}
