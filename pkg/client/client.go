package client

import (
	"context"
	"io"

	"github.com/117503445/landns/pkg/grpcgen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	target    string
	leaseChan chan<- []*grpcgen.Lease
}

func NewClient(
	target string,
	leaseChan chan<- []*grpcgen.Lease) *Client {
	return &Client{
		target:    target,
		leaseChan: leaseChan,
	}
}

func (c *Client) Start() {
	conn, err := grpc.NewClient(c.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial")
	}
	defer conn.Close()

	client := grpcgen.NewDHCPManagerClient(conn)

	stream, err := client.GetLeases(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get leases")
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Info().Msg("EOF")
			break
		}
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to receive")
		}
		log.Info().Interface("leases", resp.Leases).Msg("Received leases")
		c.leaseChan <- resp.Leases
	}
	log.Warn().Msg("Stream closed")
}
