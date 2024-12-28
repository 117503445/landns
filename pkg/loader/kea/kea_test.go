package kea_test

import (
	"testing"

	"github.com/117503445/dhcp-manager/pkg/loader/kea"
	"github.com/117503445/goutils"
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