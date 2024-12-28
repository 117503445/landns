package main

import (
	"github.com/117503445/dhcp-manager/pkg/grpcgen"
	"github.com/117503445/dhcp-manager/pkg/loader/kea"
	"github.com/117503445/dhcp-manager/pkg/server"
	"github.com/117503445/goutils"
	// "github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Hello, World!")

	leaseChan := make(chan []*grpcgen.Lease)
	go func() {
		server.NewServer(leaseChan).Start(8080)
	}()

	// go func() {
	// 	for leases := range leaseChan {
	// 		log.Info().Interface("leases", leases).Msg("Parsed leases")
	// 	}
	// }()

	err := kea.ParseStream("/var/lib/kea/dhcp4.leases", leaseChan)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse stream")
	}
}
