package cli

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

type LandnsTarget struct {
	Host string
}

var LandnsAgentCli struct {
	KeaLeaseDir string `description:"Kea lease directory" default:"/var/lib/kea/dhcp4.leases"`
	Targets     []*LandnsTarget
	Name        string
}

func LoadLandnsAgentCli() {
	cfgFile := filepath.Dir(os.Args[0]) + "/config.toml"
	// log.Info().Str("cfgFile", cfgFile).Msg("Loading config")
	kong.Parse(&LandnsAgentCli, kong.Configuration(kongtoml.Loader, cfgFile))

	if LandnsAgentCli.Name == "" {
		log.Fatal().Msg("Agent name is empty")
	}

	if len(LandnsAgentCli.Targets) == 0 {
		log.Warn().Msg("No targets specified, use default")
		LandnsAgentCli.Targets = append(LandnsAgentCli.Targets, &LandnsTarget{
			Host: "http://localhost:4500",
		})
	} else {
		for _, target := range LandnsAgentCli.Targets {
			if target.Host == "" {
				log.Fatal().Msg("Target host is empty")
			}
		}
	}

	log.Info().Interface("LandnsAgentCli", LandnsAgentCli).Send()
}
