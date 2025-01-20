package kea_test

import (
	"testing"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/loader/kea"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

func TestParse(t *testing.T) {
	content, err := goutils.ReadText("../../../assets/dhcp4.leases")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}
	leases, err := kea.Parse(content)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse leases")
	}
	log.Info().Interface("leases", leases).Msg("Parsed leases")
}

func TestStreamParse(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create watcher")
	}
	defer watcher.Close()
	err = watcher.Add("./data")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to add file to watcher")
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Debug().Interface("event", event).Msg("")

		case err, ok := <-watcher.Errors:
			if !ok {
				log.Fatal().Err(err).Msg("Failed to read file")
			}
		}
	}
}
