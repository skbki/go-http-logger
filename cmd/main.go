package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var log *logrus.Entry

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log = logrus.WithFields(logrus.Fields{
		"module": "http-logger",
	})
	log.Info("startup")
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/{rest:.*}", logPrintHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func logPrintHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(os.Stdout, r.Body)
	if r.ContentLength > 0 {
		fmt.Println()
	}
	if err != nil {
		log.Error(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
