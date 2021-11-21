package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	ClientIP := c.ClientIP()
	RemoteIP, _ := c.RemoteIP()
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if RemoteIP.String() != "127.0.0.1" {
		ip = RemoteIP.String()
	}
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}

// GetLocation 获取外网ip地址
func GetLocation(ip, key string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	url := "https://restapi.amap.com/v5/ip?ip=" + ip + "&type=4&key=" + key
	fmt.Println("url", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("restapi.amap.com failed:", err)
		return "未知位置"
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		fmt.Println("Unmarshal failed:", err)
	}
	//if m["province"] == "" {
	//	return "未知位置"
	//}
	return m["country"] + "-" + m["province"] + "-" + m["city"] + "-" + m["district"] + "-" + m["isp"]
}

// GetLocalHost 获取局域网ip地址
func GetLocalHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addr, _ := netInterfaces[i].Addrs()

			for _, address := range addr {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						return ipNet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
