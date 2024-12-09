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
		HostPolicy: autocert.HostWhitelist("hephaistos.hafa.fr"), //Your domain here
		Cache:      autocert.DirCache("./certs"),                 //Folder for storing certificates
	}
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
		GetCertificate: certManager.GetCertificate,
	}
	s := &server.Server{
		Config: http.Server{
			Addr:         ":8443",
			TLSConfig:    &tls_cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		},
	}
	s.Config.TLSConfig.NextProtos = append([]string{"h2", "http/1.1"}, s.Config.TLSConfig.NextProtos...)
	s.HandleRoutes()

	log.Println("Listening on https://localhost:8443/")
	if err := s.Config.ListenAndServeTLS("./certs/server.crt", "./certs/server.key"); err != nil {
		log.Println(err)
		log.Fatal("Can't establish connection")
	}
}
