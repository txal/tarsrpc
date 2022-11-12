package util

import (
	//	"fmt"
	"net"
)

//获取自己的ip地址列表
func GetSelfIps() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}

//判断ip是否是自己
func IsSelfIp(ip string) (ret bool, err error) {
	ips, err := GetSelfIps()
	if err != nil {
		return false, err
	}

	for _, a := range ips {
		if a == ip {
			return true, nil
		}
	}
	return false, nil
}

//获取ip地址，其中一个
func IpGetOne() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// 获取指定网络名的 v4 ip
func ipV4ByEthName(ethname string) string {
	var ip string
	ethinterface, err := net.InterfaceByName(ethname)
	if err != nil || ethinterface == nil {
		return ip
	}

	addrs, err := ethinterface.Addrs()
	if err != nil {
		return ip
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip
}

/***********
  下面函数针对多玩服务器的线上配置获取对应的ip信息
*/

// 多玩网络名配置
const (
	CTL_IP     = "eth0"   // 电信ip
	CNC_IP     = "eth0:1" // 网通ip
	MOB_IP     = "eth0:2" // 移动网络ip
	MANAGER_IP = "eth1"   // 内网管理ip
)

// 返回地址时,优先级为 电信>网通>移动
func GetDWPublicIp() string {
	var ip string
	for {
		if ip = GetDWCTLIp(); len(ip) > 0 {
			break
		}
		if ip = GetDWCNCIp(); len(ip) > 0 {
			break
		}
		if ip = GetDWMOBIp(); len(ip) > 0 {
			break
		}
		break
	}
	return ip
}

func GetDWCTLIp() string {
	return ipV4ByEthName(CTL_IP)
}

func GetDWCNCIp() string {
	return ipV4ByEthName(CNC_IP)
}

func GetDWMOBIp() string {
	return ipV4ByEthName(MOB_IP)
}

func GetDWManagerIp() string {
	return ipV4ByEthName(MANAGER_IP)
}
