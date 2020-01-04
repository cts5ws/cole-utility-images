package reprox

import (
	"fmt"
	"github.com/cts5ws/ratelim"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var protocol = "http"
var downstreamUrl = "webserver"
var port = 80

var downstreamBaseUrlRaw = fmt.Sprintf("%s://%s:%d", protocol, downstreamUrl, port)
var downstreamBaseUrl, _ = url.Parse(downstreamBaseUrlRaw)

func Route(res http.ResponseWriter, req *http.Request) {

	proxyIp, originIp := getOriginAndProxyIpAddresses(req)

	if ratelim.ShouldRateLimit(originIp) {
		http.Error(res, "Enhance your calm bro", 420)
	} else {
		fmt.Printf("Proxying to %s from (%s origin: %s) \n", req.URL.Path, proxyIp, originIp)
		forwardRequestDownstream(res, req)
	}
}

func getOriginAndProxyIpAddresses(req *http.Request) (proxyIp, originIp net.IP) {
	directRequesterIp := net.ParseIP(strings.Split(req.RemoteAddr, ":")[0])
	potentialOriginIp := net.ParseIP(req.Header.Get("X-Forwarded-For"))
	if potentialOriginIp == nil {
		return nil, directRequesterIp
	}

	return directRequesterIp, potentialOriginIp
}

func forwardRequestDownstream(res http.ResponseWriter, req *http.Request) {
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(downstreamBaseUrl)

	// Update the headers to allow for SSL redirection
	req.URL.Host = downstreamBaseUrl.Host
	req.URL.Scheme = downstreamBaseUrl.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = downstreamBaseUrl.Host

	proxy.ServeHTTP(res, req)
}
