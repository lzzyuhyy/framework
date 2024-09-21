package getipv4address

import (
	"fmt"
	"net"
)

func GetLocalIPv4() (net.IP, error) {
	var wlanIP, ethernetIP net.IP
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue // 跳过未激活和回环接口
		}

		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && ip.To4() != nil { // 只考虑IPv4地址
				fmt.Printf("%v: %v\n", i.Name, ip)
				if i.Name == "WLAN" {
					fmt.Println(1)
					wlanIP = ip
				} else if i.Name == "Ethernet" || i.Name == "以太网" {
					fmt.Println(2)
					ethernetIP = ip
				}
			}
		}
	}

	if wlanIP != nil {
		return wlanIP, nil
	}
	if ethernetIP != nil {
		return ethernetIP, nil
	}
	return nil, fmt.Errorf("%v", "No Used Public IP")
}
