package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const BASE_URL = "localhost"
const PORT = ":3333"

type Server struct {
	mappings map[string]string
	mu       sync.RWMutex
}

type Body struct {
	Endpoint    string `json:"endpoint"`
	RedirectURL string `json:"url"`
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	var body = Body{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json body.", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.mappings[body.Endpoint] = body.RedirectURL
	w.Write([]byte("added endpoint: /" + body.Endpoint + "\n"))
	fmt.Printf("\nadded '/%s' mapping to '%s'\n", body.Endpoint, s.mappings[body.Endpoint])
	s.mu.Unlock()
}

func (s *Server) redirect(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	s.mu.RLock()
	url, ok := s.mappings[endpoint]
	if !ok {
		s.mu.RUnlock()
		http.Error(w, "endpoint is not saved as a redirect.", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
	fmt.Printf("\nredirected '%s' to '%s'\n", endpoint, s.mappings[endpoint])
	s.mu.RUnlock()
}

func (s *Server) remove(w http.ResponseWriter, r *http.Request) {
	var body = Body{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json body.", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	if _, ok := s.mappings[body.Endpoint]; !ok {
		http.Error(w, "endpoint does not exist.", http.StatusNotFound)
		s.mu.Unlock()
		return
	}
	delete(s.mappings, body.Endpoint)
	w.Write([]byte("deleted endpoint.\n"))
	fmt.Printf("\nremoved %s mapping\n", body.Endpoint)
	s.mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root.\n"))
}

func main() {
	srv := &Server{mappings: map[string]string{"test": "https://tesla.com"}}
	srv.mappings["google"] = "https://google.com"

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler)
	r.Post("/", srv.create)
	r.Delete("/", srv.remove)
	r.Get("/{endpoint}", srv.redirect)

	fmt.Printf("listening on %s\n", PORT)
	http.ListenAndServe(BASE_URL+PORT, r)
}
