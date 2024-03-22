package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "192.168.0.1"
	isValidIPv4 := isIPv4(ip)

	if isValidIPv4 {
		fmt.Println(ip, "合法的IPv4地址.")
	} else {
		fmt.Println(ip, "不是合法的IPv4地址.")
	}
}

func isIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
  // 使用正则表达式
  // import "regexp"
  //ipv4Regex := `^(\d{1,3}\.){3}\d{1,3}$`
	//match, _ := regexp.MatchString(ipv4Regex, ip)
	//return match

  // 使用net.ParseCIDR函数
  //_, _, err := net.ParseCIDR(ip + "/32")
	//return err == nil
}
