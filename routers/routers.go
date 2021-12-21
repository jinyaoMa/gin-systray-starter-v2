package routers

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

var (
	handler     *gin.Engine
	httpServer  *http.Server
	httpsServer *http.Server
)

func init() {
	handler = gin.Default()
	routes(handler)

	SetConfig(DefaultConfig())
}

func SetConfig(config Config) {
	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(config.CertDirCache),
	}

	httpServer = &http.Server{
		Addr: config.HttpPort,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, config.HttpPort, config.HttpsPort, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(manager)
	httpsServer = &http.Server{
		Addr:      config.HttpsPort,
		TLSConfig: tlsConfig,
		Handler:   handler,
	}
}

func Start(callback func()) {
	go startHttpServer()
	go startHttpsServer(callback)
}

func StartWithLoop() {
	go startHttpServer()
	startHttpsServer(nil)
}

func Stop(callback func()) {
	stopHttpServer()
	stopHttpsServer(callback)
}

func startHttpsServer(callback func()) (err error) {
	err = httpsServer.ListenAndServeTLS("", "")
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return
	}
	callback()
	return nil
}

func stopHttpsServer(callback func()) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = httpsServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Https Server Shutdown:", err)
		return
	}

	log.Println("Https Server exiting!")
	callback()
	return nil
}

func startHttpServer() (err error) {
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return
	}
	return nil
}

func stopHttpServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Http Server Shutdown:", err)
		return
	}

	log.Println("Http Server exiting!")
	return nil
}

func getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = "certs"
		}

		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			log.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		log.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}
