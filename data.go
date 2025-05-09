package main

import (
	"github.com/TT-GB/nmapdata/data"
)

func GetSampleData() []data.NmapRun {
	// Initialize sample data for NmapRun
	return []data.NmapRun{
		{
			Scanner:          "nmap-v7.91",
			Args:             "-sP 192.168.1.1/24",
			Version:          "7.91",
			XMLOutputVersion: "1.04",
			Hosts: []data.Host{
				{
					StartTime: 1622538901,
					EndTime:   1622538905,
					Status: data.Status{
						State:     "up",
						Reason:    "syn-ack",
						ReasonTTL: "63",
					},
					Address: data.Address{
						Addr:     "192.168.1.1",
						AddrType: "ipv4",
					},
					Hostnames: data.Hostnames{
						Hostname: []data.Hostname{
							{Name: "router.local", Type: "PTR"},
						},
					},
					Ports: data.Ports{
						PortList: []data.Port{
							{
								Protocol: "tcp",
								PortID:   80,
								State: data.PortState{
									State:     "open",
									Reason:    "syn-ack",
									ReasonTTL: "63",
								},
								Service: &data.Service{
									Name:   "http",
									Method: "syn",
									Conf:   100,
								},
							},
						},
					},
				},
			},
			RunStats: data.RunStats{
				Finished: data.Finished{
					Time:    1622538905,
					Elapsed: "5s",
					Summary: "Nmap run completed",
					Exit:    "success",
				},
				Hosts: data.Hosts{
					Up:    1,
					Down:  0,
					Total: 1,
				},
			},
		},
		{
			Scanner:          "nmap-v7.91",
			Args:             "-sS 192.168.1.0/24",
			Version:          "7.91",
			XMLOutputVersion: "1.04",
			Hosts: []data.Host{
				{
					StartTime: 1622538910,
					EndTime:   1622538915,
					Status: data.Status{
						State:     "down",
						Reason:    "no-response",
						ReasonTTL: "64",
					},
					Address: data.Address{
						Addr:     "192.168.1.10",
						AddrType: "ipv4",
					},
					Hostnames: data.Hostnames{
						Hostname: []data.Hostname{
							{Name: "device-01.local", Type: "PTR"},
						},
					},
					Ports: data.Ports{
						PortList: []data.Port{},
					},
				},
			},
			RunStats: data.RunStats{
				Finished: data.Finished{
					Time:    1622538915,
					Elapsed: "5s",
					Summary: "Nmap run completed",
					Exit:    "success",
				},
				Hosts: data.Hosts{
					Up:    0,
					Down:  1,
					Total: 1,
				},
			},
		},
		{
			Scanner:          "nmap-v7.91",
			Args:             "-sT 192.168.2.0/24",
			Version:          "7.91",
			XMLOutputVersion: "1.04",
			Hosts: []data.Host{
				{
					StartTime: 1622538920,
					EndTime:   1622538925,
					Status: data.Status{
						State:     "up",
						Reason:    "syn-ack",
						ReasonTTL: "64",
					},
					Address: data.Address{
						Addr:     "192.168.2.2",
						AddrType: "ipv4",
					},
					Hostnames: data.Hostnames{
						Hostname: []data.Hostname{
							{Name: "device-02.local", Type: "PTR"},
						},
					},
					Ports: data.Ports{
						PortList: []data.Port{
							{
								Protocol: "tcp",
								PortID:   22,
								State: data.PortState{
									State:     "open",
									Reason:    "syn-ack",
									ReasonTTL: "64",
								},
								Service: &data.Service{
									Name:   "ssh",
									Method: "syn",
									Conf:   100,
								},
							},
						},
					},
				},
			},
			RunStats: data.RunStats{
				Finished: data.Finished{
					Time:    1622538925,
					Elapsed: "5s",
					Summary: "Nmap run completed",
					Exit:    "success",
				},
				Hosts: data.Hosts{
					Up:    1,
					Down:  0,
					Total: 1,
				},
			},
		},
	}
}
