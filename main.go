package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
)

func setupHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"msg\":\"success\"}"))
	})
	return mux
}
func setGinHandler() *gin.Engine {
	engine := gin.New()
	engine.Handle(http.MethodGet, "hi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	return engine
}

func main() {
	handler := setGinHandler()

	keyFile := "./priv.key"
	certFile := "./cert.pem"

	// visit in browser
	go func() {
		if err := http3.ListenAndServe("localhost:6121", certFile, keyFile, handler); err != nil {
			fmt.Println(err)
		}
	}()

	// udp enable
	server := http3.Server{
		Server: &http.Server{Handler: handler, Addr: "localhost:6120"},
	}

	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		fmt.Println(err)
	}
}
