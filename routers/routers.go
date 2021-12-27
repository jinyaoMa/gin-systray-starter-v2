package routers

import (
	"App/logger"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

var (
	ErrSetConfigWhileRunning error = errors.New("routers.SetConfig: cannot set config while running")
)

var (
	isRunning bool
	withSwag  bool
	config    *Config
)

var (
	server    *http.Server
	serverTls *http.Server
)

func init() {
	gin.DefaultWriter = logger.Routers.Writer()
	SetConfig(DefaultConfig())
}

func Start(swag bool, loop bool) {
	if isRunning {
		return
	}

	isRunning = true
	withSwag = swag
	handleConfig()
	if loop {
		go startServer()
		startServerTls()
	} else {
		go startServer()
		go startServerTls()
	}
}

func Stop() (ok bool) {
	if !isRunning {
		return false
	}

	isRunning = false
	stopServer()
	stopServerTls()
	return true
}

func SetConfig(conf *Config) {
	if isRunning {
		logger.Routers.Fatalln(ErrSetConfigWhileRunning)
		return
	}
	config = conf

	if config.IsDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
}

func handleConfig() {
	handler := gin.Default()
	port := fmt.Sprintf(":%d", config.Port)
	portTls := fmt.Sprintf(":%d", config.PortTls)
	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(config.CertDir),
	}
	tlsConfig := manager.TLSConfig()

	routes(handler)

	server = &http.Server{
		Addr: port,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, port, portTls, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(manager)
	serverTls = &http.Server{
		Addr:      portTls,
		TLSConfig: tlsConfig,
		Handler:   handler,
	}
}

func startServer() {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Routers.Fatalf("Server listen: %v\n", err)
		return
	}
}

func stopServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.Routers.Fatalf("Server with port %d shutdown: %v\n", config.Port, err)
		return
	}

	logger.Routers.Printf("Server with port %d is exiting!\n", config.Port)
	return
}

func startServerTls() {
	err := serverTls.ListenAndServeTLS("", "")
	if err != nil && err != http.ErrServerClosed {
		logger.Routers.Fatalf("Server (TLS) listen: %v\n", err)
		return
	}
}

func stopServerTls() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = serverTls.Shutdown(ctx)
	if err != nil {
		logger.Routers.Fatalf("Server (TLS) with port %d shutdown: %v\n", config.PortTls, err)
		return
	}

	logger.Routers.Printf("Server (TLS) with port %d is exiting!\n", config.PortTls)
	return
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
			logger.Routers.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		logger.Routers.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}
