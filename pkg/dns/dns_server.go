package dns

import (
	"fmt"
	"strings"

	"github.com/117503445/landns/pkg/store"
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

type DNSServer struct {
	udpServer  *dns.Server
	leaseStore *store.LeasesStore
}

func NewServer(leasesStore *store.LeasesStore) *DNSServer {
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		// log.Info().Msg("Got request")
		resp := new(dns.Msg)
		resp.SetReply(r)
		resp.Authoritative = true

		for _, q := range resp.Question {
			// log.Debug().Str("qname", q.Name).Str("qtype", dns.TypeToString[q.Qtype]).Msg("Question")
			// qname like archlinux.lan.
			if q.Qtype == dns.TypeA {
				if strings.HasSuffix(q.Name, ".lan.") {
					hostname := strings.TrimSuffix(q.Name, ".lan.")
					ip := leasesStore.GetIpByHostname(hostname)
					if ip != nil {
						a := &dns.A{
							Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
							A:   ip,
						}
						resp.Answer = append(resp.Answer, a)
					}
				}
			}
		}
		if len(resp.Answer) == 0 {
			resp.SetRcode(r, dns.RcodeNameError)
		}
		// m.SetRcode(r, dns.RcodeNameError)
		w.WriteMsg(resp)
	})

	return &DNSServer{
		leaseStore: leasesStore,
	}
}

func (s *DNSServer) Start(port int) {
	addr := fmt.Sprintf(":%d", port)
	udpServer := &dns.Server{Addr: addr, Net: "udp"}
	s.udpServer = udpServer

	log.Info().Int("port", port).Msg("Starting dns server")
	if err := s.udpServer.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to set udp listener")
	}
}
