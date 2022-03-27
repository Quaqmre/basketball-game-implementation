package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	gamesMutex sync.Mutex
	games      []*Game
)

func main() {
	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) { serveGames(w, r) })

	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func serveGames(w http.ResponseWriter, r *http.Request) {
	gamesMutex.Lock()
	defer gamesMutex.Unlock()

	if r.Method == http.MethodPost {
		var req []string
		err := json.NewDecoder(r.Body).Decode(&req)
		fmt.Println(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(games) == 0 {
			games = make([]*Game, len(req))
			for i, gn := range req {
				g := NewGame(gn)
				go g.Start()
				games[i] = g
			}

			w.WriteHeader(http.StatusCreated)
		}
	}

	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(games)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
