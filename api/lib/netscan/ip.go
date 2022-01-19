package netscan

import (
	"net"
	"strconv"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func getIPRange(cidr string) (addresses []string, network, broadcast, mask string, err error) {
	var ips []string
	var num int64

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	addresses = ips
	if len(ips) >= 2 {
		addresses = ips[1 : len(ips)-1]
	}
	broadcast = ips[len(ips)-1]
	network = ips[0]

	ones, bits := ipnet.Mask.Size()
	segments := 0
	bin := ""
	for i := 0; i < ones; i++ {
		bin += "1"
		if len(bin) == 8 {
			num, err = strconv.ParseInt(bin, 2, 32)
			if err != nil {
				return
			}
			mask += strconv.Itoa(int(num))
			if segments < 3 {
				mask += "."
				segments++
			}
			bin = ""
		}
	}
	for i := 0; i < (bits - ones); i++ {
		bin += "0"
		if len(bin) == 8 {
			num, err = strconv.ParseInt(bin, 2, 32)
			if err != nil {
				return
			}
			mask += strconv.Itoa(int(num))
			if segments < 3 {
				mask += "."
				segments++
			}
			bin = ""
		}
	}

	// mask = fmt.Sprintf("%d", ones)

	// fmt.Println(bits)

	return

}
