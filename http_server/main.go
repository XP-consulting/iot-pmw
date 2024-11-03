package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
    "log"
)

const (
	port         = ":8090"
	responseBody = "Welcome to Iot-pmw!"
  )

cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
if err != nil {
	log.Fatalf("Failed to load X509 key pair: %v", err)
}

func recieve_command(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "benis\n")
	fmt.Println(req)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func validate_sigML(sigml string) bool {
	//code to validate sigML
	return true

}

func main() {

	http.HandleFunc("/hello", recieve_command)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
	//http.ListenAndServe("80", nil)
}
