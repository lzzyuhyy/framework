package getipv4address

import (
	"fmt"
	"net"
)

// get ip
func GetLocalIPv4() net.IP {
	var ipAddr net.IP
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
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
			if ip == nil || ip.IsLoopback() {
				continue
			}
			// 只考虑IPv4地址
			if i.Name == "WLAN" {
				ipAddr = ip
			}
		}
	}

	return ipAddr
}
