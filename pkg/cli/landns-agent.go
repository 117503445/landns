package cli

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

var LandnsAgentCli struct {
	KeaLeaseDir string `description:"Kea lease directory" default:"/var/lib/kea/dhcp4.leases"`
}

func LoadLandnsAgentCli() {
	cfgFile := filepath.Dir(os.Args[0]) + "/config.toml"
	log.Info().Str("cfgFile", cfgFile).Msg("Loading config")
	kong.Parse(&LandnsAgentCli, kong.Configuration(kongtoml.Loader, cfgFile))
	log.Info().Interface("LandnsAgentCli", LandnsAgentCli).Send()
}
