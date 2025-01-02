package main

import (
	"context"
	"net/http"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/dns"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/rpclogic"
	"github.com/117503445/landns/pkg/store"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	goutils.InitZeroLog()

	leasesStore := store.NewLeasesStore()
	// leasesStore.SetLeases("agent1", []*rpcgen.Lease{
	// 	{
	// 		Ip:       "192.168.100.101",
	// 		Hostname: "archlinux",
	// 	}})

	go func() {
		rpclogic.NewLandnsServer(leasesStore).Start(4500)
	}()

	go func() {
		targets := []string{"http://localhost:4501"}
		log.Info().Strs("targets", targets).Msg("get leases from agents at start")

		clients := make([]rpcgen.LanDNSAgent, len(targets))
		for i, target := range targets {
			clients[i] = rpcgen.NewLanDNSAgentProtobufClient(target, &http.Client{})
		}
		for _, client := range clients {
			go func(client rpcgen.LanDNSAgent) {
				resp, err := client.GetLeases(context.Background(), &emptypb.Empty{})
				if err != nil {
					log.Error().Err(err).Msg("Failed to get leases")
					return
				}
				leasesStore.SetLeases(resp.AgentName, resp.Leases)
			}(client)
		}
	}()

	dns.NewServer(leasesStore).Start()
}
