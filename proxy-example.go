package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// targetURL digunakan untuk menentukan alamat tujuan dari proxy
	targetURL, _ := url.Parse("http://example.com")
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Menentukan alamat tujuan
		r.Host = targetURL.Host
		// Mengirimkan permintaan ke alamat tujuan
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting proxy server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// gye wang
