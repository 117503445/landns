package server

import (
	"context"
	"fmt"
	"net"

	"github.com/117503445/dhcp-manager/pkg/grpcgen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	grpcgen.DHCPManagerServer
}

func (s *Server) GetLeases(context.Context, *emptypb.Empty) (*grpcgen.GetLeasesResponse, error) {
	return &grpcgen.GetLeasesResponse{}, nil
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}
	grpcS := grpc.NewServer()
	grpcgen.RegisterDHCPManagerServer(grpcS, s)

	if err := grpcS.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
}
