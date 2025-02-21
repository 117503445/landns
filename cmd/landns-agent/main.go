package main

import (
	"context"
	"net/http"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/cli"
	"github.com/117503445/landns/pkg/loader/kea"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/rpclogic"
	"github.com/117503445/landns/pkg/util"

	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()

	cli.LoadLandnsAgentCli()

	agentName := cli.LandnsAgentCli.Name

	leaseChan := make(chan []*rpcgen.Lease)
	agentServer := rpclogic.NewLandnsAgentServer(agentName)

	go func() {
		agentServer.Start(4501)
	}()

	go func() {

		clients := make([]rpcgen.LanDNS, len(cli.LandnsAgentCli.Targets))
		for i, target := range cli.LandnsAgentCli.Targets {
			clients[i] = rpcgen.NewLanDNSProtobufClient(target.Host, &http.Client{})
		}

		executors := make([]*util.LatestTaskExecutor, len(clients))
		for i := range clients {
			executors[i] = util.NewLatestTaskExecutor()
			executors[i].Start()
		}

		// leaseChan -> clients
		for leases := range leaseChan {
			log.Info().Interface("leases", leases).Msg("leaseChan -> broadcast to clients")
			agentServer.SetLeases(leases)

			for i, client := range clients {
				executors[i].AddTask(func() {
					log.Debug().Str("target", cli.LandnsAgentCli.Targets[i].Host).Msg("Sending leases")
					_, err := client.SetLeases(context.Background(), &rpcgen.SetLeasesRequest{
						Leases:    leases,
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
	err := kea.ParseStream(cli.LandnsAgentCli.KeaLeaseDir, leaseChan)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse stream")
	}
}
