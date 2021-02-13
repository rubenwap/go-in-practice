package main

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
		Dial: (&net.Dialer{Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://www.example.com")
	if err != nil {
		//
	}
	o, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//
	}
	resp.Body.Close()
	// If body is closed, next request can reuse the connection kept alive
}
