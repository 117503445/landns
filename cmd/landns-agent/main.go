package main

import (
	"context"
	"net/http"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/loader/kea"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/117503445/landns/pkg/rpclogic"

	// "github.com/117503445/landns/pkg/server"

	// "github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Hello, World!")

	leaseChan := make(chan []*rpcgen.Lease)
	go func() {
		rpclogic.NewLandnsAgentServer().Start(4501)
	}()

	go func() {
		targets := []string{"http://localhost:4500"}

		clients := make([]rpcgen.LanDNS, len(targets))
		for i, target := range targets {
			clients[i] = rpcgen.NewLanDNSProtobufClient(target, &http.Client{})
		}

		chans := make([]chan []*rpcgen.Lease, len(clients))
		for i, client := range clients {
			chans[i] = make(chan []*rpcgen.Lease, 10)
			go func(client rpcgen.LanDNS, leaseChan <-chan []*rpcgen.Lease) {
				leases := <-leaseChan
				// 如果 leaseChan 还有数据，就继续读取
				for {
					isBreak := false
					select {
					case leases = <-leaseChan:
						log.Info().Interface("leases", leases).Msg("Replace leases")
					default:
						isBreak = true
						break
					}
					if isBreak {
						break
					}
				}

				_, err := client.SetLeases(context.Background(), &rpcgen.SetLeasesRequest{
					Leases: <-leaseChan,
				})
				if err != nil {
					log.Error().Err(err).Msg("Failed to set leases")
				}
			}(client, chans[i])
		}

		for leases := range leaseChan {
			log.Info().Interface("leases", leases).Msg("Parsed leases")

		}
	}()

	err := kea.ParseStream("/var/lib/kea/dhcp4.leases", leaseChan)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse stream")
	}
}
