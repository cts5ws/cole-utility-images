package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
)

var protocol = "http"
var downstreamUrl = "webserver"
var port = 80

var downstreamBaseUrlRaw = fmt.Sprintf("%s://%s:%d", protocol, downstreamUrl, port)

func routeToLocalPort(res http.ResponseWriter, req *http.Request) {

    proxyIp := strings.Split(req.RemoteAddr, ":")[0]
    originIp := req.Header.Get("'X-Forwarded-For")
    if originIp == "" {
        originIp = "n/a"
    }

    fmt.Printf("Proxying to %s from (%s origin: %s) \n", req.URL.Path, proxyIp, originIp)

    downstreamBaseUrl, _ := url.Parse(downstreamBaseUrlRaw)

    // create the reverse proxy
    proxy := httputil.NewSingleHostReverseProxy(downstreamBaseUrl)

    // Update the headers to allow for SSL redirection
    req.URL.Host = downstreamBaseUrl.Host
    req.URL.Scheme = downstreamBaseUrl.Scheme
    req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
    req.Host = downstreamBaseUrl.Host

    proxy.ServeHTTP(res, req)
}

func main() {

    http.HandleFunc("/", routeToLocalPort)

    http.ListenAndServe(":80", nil)
}
