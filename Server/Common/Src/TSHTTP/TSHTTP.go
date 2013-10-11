package TSHTTP

import (
	"net/http"
	"log"
)

type ReceiveBuffer func (w http.ResponseWriter, r *http.Request)
type ServerInit func()

func CreateHTTPServer(sWebPath string, si ServerInit, rb ReceiveBuffer) {
	http.HandleFunc("/", rb)
	si();
	log.Fatal(http.ListenAndServe(sWebPath, nil));
}
