package ratelim

import (
	"net"
	"time"
)

var requestWaitTime, _ = time.ParseDuration("10s")

var ipAddrToLastRequest = make(map[string]time.Time)

func ShouldRateLimit(ip net.IP) bool {

	currentTime := time.Now()

	lastRequestTime := ipAddrToLastRequest[ip.String()]

	if currentTime.Sub(lastRequestTime) > requestWaitTime {
		ipAddrToLastRequest[ip.String()] = currentTime
		return false
	} else {
		return true
	}
}
