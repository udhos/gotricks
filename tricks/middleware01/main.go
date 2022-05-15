package main

import (
	"log"
	"net/http"
)

func main() {

	listenAddr := ":8080"

	mux := http.NewServeMux()

	log.Printf("installing handleRoot at /")
	mux.HandleFunc("/", logger(basicAuth(handleRoot)))

	log.Printf("installing handleHello at /hello")
	mux.HandleFunc("/hello", logger(basicAuth(handleHello)))

	log.Printf("server listening at %s", listenAddr)
	server := &http.Server{Addr: listenAddr, Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func logger(h http.HandlerFunc) http.HandlerFunc {
	const me = "logger"
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: method=%s host=%s path=%s from=%s", me, r.Method, r.Host, r.URL.Path, r.RemoteAddr)
		h(w, r)
	}
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	const me = "basicAuth"
	return func(w http.ResponseWriter, r *http.Request) {

		username, password, authOK := r.BasicAuth()
		if !authOK {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			log.Printf("%s url=%s from=%s Basic Auth missing", me, r.URL.Path, r.RemoteAddr)
			http.Error(w, "401 Unauthenticated", 401)
			return
		}

		if username != "admin" || password != "admin" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			log.Printf("%s url=%s from=%s Basic Auth failed", me, r.URL.Path, r.RemoteAddr)
			http.Error(w, "401 Unauthenticated", 401)
			return
		}

		log.Printf("%s url=%s from=%s Basic Auth ok", me, r.URL.Path, r.RemoteAddr)

		h(w, r)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	const me = "handleRoot"
	const notFound = "404 nothing here"
	log.Printf("%s: method=%s host=%s path=%s from=%s - response:%s", me, r.Method, r.Host, r.URL.Path, r.RemoteAddr, notFound)
	http.Error(w, notFound, 404)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	const me = "handleHello"
	const ok = "200 ok"
	log.Printf("%s: method=%s host=%s path=%s from=%s - response:%s", me, r.Method, r.Host, r.URL.Path, r.RemoteAddr, ok)
	http.Error(w, ok, 200)
}
