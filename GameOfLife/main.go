package main

import (
	"encoding/json"
	"github.com/maxxxlounge/interviews/GameOfLife/game"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	l := log.New()

	tickTime := 1 * time.Second
	g := game.NewGame(100, 100)
	g.Generate()
	go func(g *game.Game) {
		for {
			g.Tick()
			time.Sleep(tickTime)
		}
	}(g)

	var h http.Handler

	http.HandleFunc("/cells", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		out, err := json.Marshal(g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(out)
	})

	http.HandleFunc("/cells/generate", func(w http.ResponseWriter, r *http.Request) {
		g.Generate()
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	s := &http.Server{
		Addr:           ":8888",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	l.Infof("starting endpoint on port :8888")
	l.Fatal(s.ListenAndServe())
}

