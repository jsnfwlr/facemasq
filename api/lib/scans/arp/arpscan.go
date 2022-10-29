// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// arpscan implements ARP scanning of all interfaces' local networks using
// gopacket and its subpackages.  This example shows, among other things:
//   - Generating and sending packet data
//   - Reading in packet data and interpreting it
//   - Use of the 'pcap' subpackage for reading/writing
package arpscan

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/logging"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var Frequency time.Duration

const UseScanLog = false

func init() {
	var err error
	Frequency, err = time.ParseDuration(os.Getenv("ARPSCAN_FREQUENCY"))
	if err != nil {
		Frequency = time.Duration(60) * time.Second
	}
}

func Schedule() {
	go ScanAndStore()
}

func ScanAndStore() {
	// Get a list of all interfaces.
	netFaces, err := net.Interfaces()
	if err != nil {
		logging.Errorf("ArpScan: %+v\n", err.Error())
		return
	}
	var wg sync.WaitGroup
	for _, netFace := range netFaces {
		if !strings.Contains(netFace.Name, "veth") && !strings.Contains(netFace.Name, "lo") && !strings.Contains(netFace.Name, "br-") && !strings.Contains(netFace.Name, "docker0") {
			addresses, err := netFace.Addrs()
			if err != nil {
				logging.Errorf("ArpScan: %+v\n", err.Error())
				continue
			}
			if len(addresses) > 0 {
				wg.Add(1)
				// Start up a scan on each interface.
				go func(iface net.Interface) {
					defer wg.Done()
					if err := ScanARP(&iface); err != nil {
						logging.Errorf("interface %v: %v", iface.Name, err)
					}
				}(netFace)
			}
		}
	}
	// Wait for all interfaces' scans to complete.  They'll try to run forever, but will stop on an error, so if we get past this Wait it means all attempts to write have failed.
	wg.Wait()
}

// scan scans an individual interface's local network for machines using ARP requests/replies.
//
// scan loops forever, sending packets out regularly.  It returns an error if
// it's ever unable to write a packet.
func ScanARP(iface *net.Interface) error {
	// We just look for IPv4 addresses, so try to find if the interface has one.
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	// Sanity-check that the interface has a good address.
	if addr == nil {
		return errors.New("no good IP network found")
	} else if addr.IP[0] == 127 {
		return errors.New("skipping localhost")
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return errors.New("mask means network is too large")
	}

	// Open up a pcap handle for packet reads/writes.
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Start up a goroutine to read in packet data.
	stop := make(chan struct{})
	go readARP(handle, iface, stop)
	defer close(stop)
	for {
		// Write our scan packets out to the handle.
		if err := writeARP(handle, iface, addr); err != nil {
			logging.Errorf("error writing packets on %v: %v", iface.Name, err)
			return err
		}
		// We don't know exactly how long it'll take for packets to be
		// sent back to us, but 10 seconds should be more than enough
		// time ;)
		time.Sleep(Frequency)
	}
}

// readARP watches a handle for incoming ARP responses we might care about, and prints them.
//
// readARP loops until 'stop' is closed.
func readARP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			logging.Processln("Stopping ARP Scan")
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				// This is a packet I sent.
				continue
			}
			// Note:  we might get some packets here that aren't responses to ones we've sent,
			// if for example someone else sends US an ARP request.  Doesn't much matter, though...
			// all information is good information :)
			var scanID int64
			var sqlres sql.Result

			err := db.Conn.NewRaw(fmt.Sprintf(`SELECT ID FROM Scans WHERE Time LIKE '%s%%'`, time.Now().Format("2006-01-02 15:04"))).Scan(db.Context, &scanID)
			if err != nil {
				if err.Error() == "sql: no rows in result set" {
					sql := `INSERT INTO Scans (Time) VALUES (?);`
					sqlres, err = db.Conn.Exec(sql, time.Now().Format("2006-01-02 15:04:05"))
					if err != nil {
						logging.Errorf("api/lib/arpscan/arpscan.go:177 Err: %v\n", err)
					}
					scanID, _ = sqlres.LastInsertId()
				} else {
					logging.Errorf("api/lib/arpscan/arpscan.go:181 Err: %v\n", err)
				}
			}

			record := Result{
				ScanID:   scanID,
				LastSeen: time.Now().Format("2006-01-02 15:04:05"),
				IPv4:     net.IP(arp.SourceProtAddress).String(),
				MAC:      strings.ToUpper(net.HardwareAddr(arp.SourceHwAddress).String()),
			}
			record.Store()
			logging.Printf(1, "Updating %v (%v) via ARP\n", record.IPv4, record.MAC)
		}
	}
}

// writeARP writes an ARP request for each address on our local network to the
// pcap handle.
func writeARP(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) (err error) {
	// Set up all the layers' fields we can.
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	// Set up buffer and options for serialization.
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	// Send one packet for every address.
	for _, ip := range ips(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
		if err = handle.WritePacketData(buf.Bytes()); err != nil {
			return
		}
	}
	return
}

// ips is a simple and not very good method for getting all IPv4 addresses from a
// net.IPNet.  It returns all IPs it can over the channel it sends back, closing
// the channel when done.
func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}
