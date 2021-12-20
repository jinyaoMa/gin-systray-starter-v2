package routers

import (
	"context"
	"log"
	"net/http"
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

func SetConfig(config *Config) {
	manager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(config.CertDirCache),
	}

	httpServer = &http.Server{
		Addr:    config.HttpPort,
		Handler: manager.HTTPHandler(http.HandlerFunc(redirect)),
	}
	httpServer.RegisterOnShutdown(onStop)

	httpsServer = &http.Server{
		Addr:      config.HttpsPort,
		TLSConfig: manager.TLSConfig(),
		Handler:   handler,
	}
}

func Start() {
	go startHttp()
	go startHttps()
}

func StartWithLoop() {
	go startHttp()
	startHttps()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Http Server Shutdown:", err)
	}

	log.Println("Http Server exiting!")
}

func onStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := httpsServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Https Server Shutdown:", err)
	}

	log.Println("Https Server exiting!")
}

func startHttps() {
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func startHttp() {
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.RequestURI

	http.Redirect(w, req, target, http.StatusMovedPermanently)
}
