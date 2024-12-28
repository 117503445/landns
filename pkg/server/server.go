package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/117503445/dhcp-manager/pkg/grpcgen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	grpcgen.DHCPManagerServer
	leaseChan <-chan []*grpcgen.Lease

	streams     []grpc.ServerStream
	streamsLock sync.RWMutex
}

func (s *Server) GetLeases(_ *emptypb.Empty, stream grpc.ServerStreamingServer[grpcgen.GetLeasesResponse]) error {
	s.streamsLock.Lock()
	defer s.streamsLock.Unlock()
	s.streams = append(s.streams, stream)
	return nil
}

func NewServer(leaseChan <-chan []*grpcgen.Lease) *Server {
	return &Server{
		leaseChan: leaseChan,
		streams:   make([]grpc.ServerStream, 0),
	}
}

func (s *Server) Start(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}
	grpcS := grpc.NewServer()
	grpcgen.RegisterDHCPManagerServer(grpcS, s)

	go func() {
		broadcast := func(leases []*grpcgen.Lease) {
			var wg sync.WaitGroup

			failedStreams := make(map[grpc.ServerStream]struct{})

			s.streamsLock.RLock()
			defer s.streamsLock.RUnlock()

			for _, stream := range s.streams {
				wg.Add(1)
				go func(stream grpc.ServerStream) {
					defer wg.Done()
					// TODO: timeout
					if err := stream.SendMsg(&grpcgen.GetLeasesResponse{Leases: leases}); err != nil {
						log.Error().Err(err).Msg("Failed to send message")
						failedStreams[stream] = struct{}{}
					}
				}(stream)
			}
			wg.Wait()

			if len(failedStreams) > 0 {
				s.streamsLock.Lock()
				defer s.streamsLock.Unlock()

				newStreams := make([]grpc.ServerStream, 0, len(s.streams)-len(failedStreams))
				for _, stream := range s.streams {
					if _, ok := failedStreams[stream]; !ok {
						newStreams = append(newStreams, stream)
					}
				}
				s.streams = newStreams
			}
		}

		for leases := range s.leaseChan {
			broadcast(leases)
		}
	}()
	if err := grpcS.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
}
