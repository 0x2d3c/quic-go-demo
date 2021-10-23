package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

func main() {

	// os: windows
	pool := x509.NewCertPool()
	// os: linux
	// pool, err := x509.SystemCertPool()

	//caBytes, err := ioutil.ReadFile("ca.pem")
	//if err != nil {
	//	panic(err)
	//}
	//if ok := pool.AppendCertsFromPEM(caBytes); !ok {
	//	panic("add ca to pool fail.")
	//}
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
		},
	}
	defer roundTripper.Close()

	hclient := &http.Client{
		Transport: roundTripper,
	}

	//rsp, err := hclient.Get("https://localhost:6121/hi")
	rsp, err := hclient.Get("https://localhost:6120/hi")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	fmt.Printf("body: %#v\n", rsp)

	body := &bytes.Buffer{}
	if _, err := body.ReadFrom(rsp.Body); err != nil {
		log.Fatal(err)
	}
	fmt.Println(body.String())
}
