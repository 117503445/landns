package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/client"
	"github.com/117503445/landns/pkg/dns"
	"github.com/117503445/landns/pkg/grpcgen"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/rs/zerolog/log"
)



func main() {
	goutils.InitZeroLog()

	log.Info().Msg("localdns start")

	leaseChan := make(chan []*rpcgen.Lease)

	go func() {
		c := client.NewClient("localhost:4358", leaseChan)
		c.Start()
	}()

	dns.NewServer().Start()
}
