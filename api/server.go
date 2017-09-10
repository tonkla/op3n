package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	truefx "github.com/tonkla/gotruefx"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", home)
	r.Route("/forex", func(r chi.Router) {
		r.Get("/", getForexAll)
		r.Get("/{symbol}", getForexBySymbol)
	})

	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func getForexAll(w http.ResponseWriter, r *http.Request) {
	result := truefx.NewFeed().Get()
	w.Header().Set("Content-Type", "application/json")
	if len(result) > 0 {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode([]string{})
	}
}

func getForexBySymbol(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	result := truefx.NewFeed().GetBySymbol(symbol)
	w.Header().Set("Content-Type", "application/json")
	if len(result) > 0 {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode([]string{})
	}
}
