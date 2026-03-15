package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	const fileServerPrefix = "/app/"

	fileServer := http.FileServer(http.Dir(filepathRoot))
	mux := http.NewServeMux()
	mux.Handle(fileServerPrefix, http.StripPrefix(fileServerPrefix, fileServer))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		okStr := "OK"
		w.Write([]byte(okStr))
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

type HealthHandler struct{}

func (h *HealthHandler) ServerHttp(writer http.ResponseWriter, req *http.Request) {

}
