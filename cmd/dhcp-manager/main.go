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

	go func() {
		server.NewServer().Start(8080)
	}()

	leaseChan := make(chan []*grpcgen.Lease)

	go func() {
		for leases := range leaseChan {
			log.Info().Interface("leases", leases).Msg("Parsed leases")
		}
	}()

	err := kea.ParseStream("/var/lib/kea/dhcp4.leases", leaseChan)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse stream")
	}

	// // Create new watcher.
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to create watcher")
	// }
	// defer watcher.Close()

	// // Start listening for events.
	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Debug().Interface("event", event).Msg("event")
	// 			if event.Has(fsnotify.Write) {
	// 				log.Debug().Str("event", "modified").Str("file", event.Name).Msg("file modified")
	// 			}
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Error().Err(err).Msg("error")
	// 		}
	// 	}
	// }()

	// // Add a path.
	// err = watcher.Add("/tmp")
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to add path")
	// }

	// // Block main goroutine forever.
	// <-make(chan struct{})

}
