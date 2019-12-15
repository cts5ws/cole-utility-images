package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
)
// ToDo:
// - add better logging
// - do implement rate limiting

func routeToLocalPort(res http.ResponseWriter, req *http.Request) {

    url, _ := url.Parse("http://webserver:80")

    // create the reverse proxy
    proxy := httputil.NewSingleHostReverseProxy(url)

    // Update the headers to allow for SSL redirection
    req.URL.Host = url.Host
    req.URL.Scheme = url.Scheme
    req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
    req.Host = url.Host

    fmt.Println("Proxying to webserver")

    proxy.ServeHTTP(res, req)
}

func main() {

    http.HandleFunc("/", routeToLocalPort)

    http.ListenAndServe(":80", nil)
}
