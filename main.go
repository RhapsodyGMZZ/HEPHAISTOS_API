package main

import (
	"crypto/tls"
	"hephaistos/server"
	"log"
	"net/http"
)

func main() {
	tls_cfg := tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	s := &server.Server{
		Config: http.Server{
			Addr:         ":8443",
			TLSConfig:    &tls_cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		},
	}
	s.HandleRoutes()

	log.Println("Listening on https://localhost:8443/")
	//	if err := s.Config.ListenAndServeTLS("/etc/letsencrypt/live/hephaistos.hafa.fr/fullchain.pem", "/etc/letsencrypt/live/hephaistos.hafa.fr/privkey.pem"); err != nil {
	if err := s.Config.ListenAndServeTLS("./certs/server.crt", "./certs/server.key"); err != nil {
		log.Println(err)
		log.Fatal("Can't establish connection")
	}
}
