package getipv4address

import (
	"fmt"
	"net"
)

// get ip
func GetLocalIPv4() (ip string, err error) {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ip, err
	}
	for _, addr := range addrList {
		// 过滤掉回环地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			// 如果ip不符合ipv4格式，继续查找下一个
			if ipv4 == nil {
				continue
			}
			return ipv4.String(), nil
		}
	}
	return ip, fmt.Errorf("not find ipv4 addr")
}
