package cli

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

type AgentTarget struct {
	Host string
}

var LandnsCli struct {
	Targets []*AgentTarget
}

func LoadLandnsCli() {
	cfgFile := filepath.Dir(os.Args[0]) + "/config.toml"
	// log.Info().Str("cfgFile", cfgFile).Msg("Loading config")
	kong.Parse(&LandnsCli, kong.Configuration(kongtoml.Loader, cfgFile))

	if len(LandnsCli.Targets) == 0 {
		log.Warn().Msg("No targets specified, use default")
		LandnsCli.Targets = append(LandnsCli.Targets, &AgentTarget{
			Host: "http://localhost:4501",
		})
	}else{
		for _, target := range LandnsCli.Targets {
			if target.Host == "" {
				log.Fatal().Msg("Target host is empty")
			}
		}
	}

	log.Info().Interface("LandnsCli", LandnsCli).Send()
}
