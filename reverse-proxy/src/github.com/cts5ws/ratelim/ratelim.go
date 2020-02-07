package ratelim

import (
	"net"
	"time"
)

var ratelimEnable = false;
var requestWaitTime, _ = time.ParseDuration("10s")

var ipAddrToLastRequest = make(map[string]time.Time)

// todo: lock map access. might not be thread safe
func ShouldRateLimit(ip net.IP) bool {

	if !ratelimEnable {
		return false
	}

	currentTime := time.Now()

	lastRequestTime := ipAddrToLastRequest[ip.String()]

	if currentTime.Sub(lastRequestTime) > requestWaitTime {
		ipAddrToLastRequest[ip.String()] = currentTime
		return false
	} else {
		return true
	}
}
