package main

import (
	"crypto/tls"
	"hephaistos/server"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("hephaistos.hafa.fr"),
		Cache:      autocert.DirCache("certs"),
	}
	tls_cfg := tls.Config{
		GetCertificate: certManager.GetCertificate,
	}
	s := &server.Server{
		Config: http.Server{
			Addr:         ":8443",
			TLSConfig:    &tls_cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		},
	}
	s.HandleRoutes()
	go func() {
		if err := http.ListenAndServe(":http", certManager.HTTPHandler(nil)); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Listening on https://localhost:8443/")
	if err := s.Config.ListenAndServeTLS("", ""); err != nil {
		log.Println(err)
		log.Fatal("Can't establish connection")
	}
}
