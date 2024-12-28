package main

import (
	"github.com/117503445/goutils"
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	log.Info().Msg("Got request")
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	for _, q := range m.Question {
		if q.Qtype == dns.TypeA {
			rr, err := dns.NewRR("example.org. 60 IN A 127.0.0.1")
			if err != nil {
				log.Printf("Error creating RR: %v", err)
				return
			}
			m.Answer = append(m.Answer, rr)
		}
	}

	w.WriteMsg(m)
}

func main() {
	goutils.InitZeroLog()

	log.Info().Msg("localdns start")

	// Create a new UDP server on port 53
	udpServer := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", handleRequest)

	if err := udpServer.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to set udp listener")
	}
}
