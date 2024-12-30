package main

import (
	"context"
	"net/http"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/loader/kea"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/rpclogic"
	"github.com/117503445/landns/pkg/util"

	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Landns agent started")

	agentName := "agent1"

	leaseChan := make(chan []*rpcgen.Lease)

	agentServer := rpclogic.NewLandnsAgentServer(agentName)

	go func() {
		log.Info().Msg("Starting landns agent server on port 4501")
		agentServer.Start(4501)
	}()

	go func() {
		targets := []string{"http://localhost:4500"}

		clients := make([]rpcgen.LanDNS, len(targets))
		for i, target := range targets {
			clients[i] = rpcgen.NewLanDNSProtobufClient(target, &http.Client{})
		}

		executors := make([]*util.LatestTaskExecutor, len(clients))
		for i := range clients {
			executors[i] = util.NewLatestTaskExecutor()
			executors[i].Start()
		}

		// leaseChan -> clients
		for leases := range leaseChan {
			log.Info().Interface("leases", leases).Msg("Parsed leases")
			agentServer.SetLeases(leases)

			for i, client := range clients {
				executors[i].AddTask(func() {
					_, err := client.SetLeases(context.Background(), &rpcgen.SetLeasesRequest{
						Leases: <-leaseChan,
						AgentName: agentName,
					})
					if err != nil {
						log.Error().Err(err).Msg("Failed to set leases")
					}
				})
			}
		}
	}()

	// file changed -> leaseChan
	err := kea.ParseStream("/var/lib/kea/dhcp4.leases", leaseChan)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse stream")
	}
}
