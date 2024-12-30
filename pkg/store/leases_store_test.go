package store

import (
	"reflect"
	"testing"

	rpcgen "github.com/117503445/landns/pkg/rpcgen"
)

// TestMergeLeases tests the MergeLeases function with various scenarios.
func TestMergeLeases(t *testing.T) {
	tests := []struct {
		name            string
		leasesByTag     map[string][]*rpcgen.Lease
		wantHostnameIp  map[string]string
		wantRepeatLease map[*rpcgen.Lease]string
	}{
		{
			name: "no_duplicates",
			leasesByTag: map[string][]*rpcgen.Lease{
				"tag1": {&rpcgen.Lease{Mac: "mac1", Ip: "192.168.1.1", Hostname: "host1"}},
				"tag2": {&rpcgen.Lease{Mac: "mac2", Ip: "192.168.1.2", Hostname: "host2"}},
			},
			wantHostnameIp: map[string]string{
				"host1": "192.168.1.1",
				"host2": "192.168.1.2",
			},
			wantRepeatLease: make(map[*rpcgen.Lease]string),
		},
		{
			name: "duplicate_hostname_different_ip",
			leasesByTag: map[string][]*rpcgen.Lease{
				"tag1": {&rpcgen.Lease{Mac: "mac1", Ip: "192.168.1.1", Hostname: "host1"}},
				"tag2": {&rpcgen.Lease{Mac: "mac2", Ip: "192.168.1.2", Hostname: "host1"}},
			},
			wantHostnameIp: map[string]string{},
			wantRepeatLease: map[*rpcgen.Lease]string{
				{Mac: "mac1", Ip: "192.168.1.1", Hostname: "host1"}: "tag2",
				{Mac: "mac2", Ip: "192.168.1.2", Hostname: "host1"}: "tag2",
			},
		},
		{
			name: "duplicate_hostname_same_ip",
			leasesByTag: map[string][]*rpcgen.Lease{
				"tag1": {&rpcgen.Lease{Mac: "mac1", Ip: "192.168.1.1", Hostname: "host1"}},
				"tag2": {&rpcgen.Lease{Mac: "mac2", Ip: "192.168.1.1", Hostname: "host1"}},
			},
			wantHostnameIp: map[string]string{
				"host1": "192.168.1.1",
			},
			wantRepeatLease: map[*rpcgen.Lease]string{},
		},
		{
			name:            "empty_input",
			leasesByTag:     map[string][]*rpcgen.Lease{},
			wantHostnameIp:  map[string]string{},
			wantRepeatLease: make(map[*rpcgen.Lease]string),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHostnameIp, gotRepeatLease := mergeLeases(tt.leasesByTag)

			if !reflect.DeepEqual(gotHostnameIp, tt.wantHostnameIp) {
				t.Errorf("MergeLeases() gotHostnameIp = %v, want %v", gotHostnameIp, tt.wantHostnameIp)
			}

			if !compareLeaseMaps(gotRepeatLease, tt.wantRepeatLease) {
				t.Errorf("MergeLeases() gotRepeatLease = %v, want %v", gotRepeatLease, tt.wantRepeatLease)
			}
		})
	}
}

func compareLeaseMaps(got, want map[*rpcgen.Lease]string) bool {
	if len(got) != len(want) {
		return false
	}
	for k, v := range got {
		found := false
		for wk, wv := range want {
			if k.Hostname == wk.Hostname && k.Ip == wk.Ip && k.Mac == wk.Mac && v == wv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
