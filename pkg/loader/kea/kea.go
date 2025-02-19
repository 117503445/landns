package kea

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"

	"github.com/117503445/goutils"
	"github.com/117503445/landns/pkg/rpcgen"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

// Parse parses the content of a Kea lease file
// example content: see assets/dhcp4.leases
func Parse(content string) ([]*rpcgen.Lease, error) {
	reader := csv.NewReader(strings.NewReader(content))

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	ipIndex, hostnameIndex, macIndex := -1, -1, -1
	for i, v := range header {
		if v == "address" {
			ipIndex = i
		} else if v == "hostname" {
			hostnameIndex = i
		} else if v == "hwaddr" {
			macIndex = i
		}
	}
	if ipIndex == -1 {
		return nil, errors.New("address not found")
	} else if hostnameIndex == -1 {
		return nil, errors.New("hostname not found")
	} else if macIndex == -1 {
		return nil, errors.New("hwaddr not found")
	}

	var leases []*rpcgen.Lease
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		hostName := record[hostnameIndex]
		if hostName == "" || hostName == "." {
			continue
		}
		hostName = strings.TrimSuffix(hostName, ".")

		leases = append(leases, &rpcgen.Lease{
			Ip:       record[ipIndex],
			Hostname: hostName,
			Mac:      record[macIndex],
		})
	}

	// only keep the latest lease for each mac
	macMap := make(map[string]*rpcgen.Lease)
	for _, lease := range leases {
		macMap[lease.Mac] = lease
	}
	leases = make([]*rpcgen.Lease, 0, len(macMap))
	for _, lease := range macMap {
		leases = append(leases, lease)
	}

	return leases, nil
}

// ParseStream watches the Kea lease file and sends the parsed leases to the leaseChan
func ParseStream(dir string, leaseChan chan<- []*rpcgen.Lease) error {
	tryParse := func() error {
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}
		leases := make([]*rpcgen.Lease, 0)

		for _, file := range files {
			if file.IsDir() {
				log.Fatal().Interface("file", file).Msg("sub dir should not in lease dir")
			}

			fullFilePath := dir + "/" + file.Name()

			content, err := goutils.ReadText(fullFilePath)
			if err != nil {
				return err
			}
			curLeases, err := Parse(content)
			if err != nil {
				return err
			}
			leases = append(leases, curLeases...)
		}

		log.Info().Interface("leases", leases).Msg("file changed -> leaseChan")
		leaseChan <- leases
		return nil
	}
	if err := tryParse(); err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if event.Has(fsnotify.Write) {
				if err := tryParse(); err != nil {
					return err
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
		}
	}
}
