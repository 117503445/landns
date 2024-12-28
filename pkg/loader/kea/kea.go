package kea

import (
	"encoding/csv"
	"errors"
	"strings"

	"github.com/117503445/dhcp-manager/pkg/grpcgen"
	"github.com/117503445/goutils"
	"github.com/fsnotify/fsnotify"
)

// Parse parses the content of a Kea lease file
// example content: see assets/dhcp4.leases
func Parse(content string) ([]*grpcgen.Lease, error) {
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

	var leases []*grpcgen.Lease
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		leases = append(leases, &grpcgen.Lease{
			Ip:       record[ipIndex],
			Hostname: record[hostnameIndex],
			Mac:      record[macIndex],
		})
	}

	// only keep the latest lease for each mac
	macMap := make(map[string]*grpcgen.Lease)
	for _, lease := range leases {
		macMap[lease.Mac] = lease
	}
	leases = make([]*grpcgen.Lease, 0, len(macMap))
	for _, lease := range macMap {
		leases = append(leases, lease)
	}

	return leases, nil
}

// ParseStream watches the Kea lease file and sends the parsed leases to the leaseChan
func ParseStream(fileName string, leaseChan chan<- []*grpcgen.Lease) error {
	tryParse := func() error {

		content, err := goutils.ReadText(fileName)
		if err != nil {
			return err
		}
		leases, err := Parse(content)
		if err != nil {
			return err
		}
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

	err = watcher.Add(fileName)
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
