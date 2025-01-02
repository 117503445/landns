package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/117503445/goutils"
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()

	var cli struct {
		// "landns" or "landns-agent"
		Bin string `name:"bin" required:"" enum:"landns,landns-agent" desc:"The binary to build and deploy."`
	}
	kong.Parse(&cli)

	bin := cli.Bin
	log.Info().Str("bin", bin).Msg("")
	target := "root@192.168.100.101"

	dirRoot, err := goutils.FindGitRepoRoot()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get git root dir")
	}
	goutils.ExecOpt.Cwd = dirRoot

	// CGO_ENABLED=0 landns-dev go build -o landns ./cmd/landns/main.go
	// goutils.Exec(fmt.Sprintf("go build -o %v ./cmd/%v/main.go", bin, bin), goutils.WithEnv{
	// 	"CGO_ENABLED": "0",
	// })
	goutils.Exec(fmt.Sprintf("ssh %v pkill %v", target, bin), goutils.WithExecutedHandlerSlient{})
	goutils.Exec(fmt.Sprintf("ssh %v rm -f /tmp/%v", target, bin), goutils.WithExecutedHandlerSlient{})
	goutils.Exec(fmt.Sprintf("scp %v %v:/tmp", bin, target))
	// goutils.Exec(fmt.Sprintf("ssh %v /tmp/%v", target, bin))

	cmd := exec.Command("ssh", target, fmt.Sprintf("/tmp/%v", bin))
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}
