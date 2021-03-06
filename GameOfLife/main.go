package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/maxxxlounge/interviews/GameOfLife/game"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	l := log.New()

	tickTime := 2 * time.Second
	g := game.NewGame(100, 100)

	err := g.Generate()
	if err != nil {
		err = errors.Wrap(err, "error generating first grid")
		l.Fatal(err)
	}

	go func(g *game.Game) {
		for {
			g.Tick()
			time.Sleep(tickTime)
		}
	}(g)

	var h http.Handler

	// return the cells json
	http.HandleFunc("/cells", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "can't get resource using not GET method", http.StatusBadRequest)

			return
		}
		w.Header().Add("Content-Type", "application/json")
		out, err := json.Marshal(g)
		if err != nil {
			l.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
		_, err = w.Write(out)
		if err != nil {
			l.Error(err)
		}
	})

	// generate new grid with width height given
	http.HandleFunc("/cells/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "can't change resorce using not POST method", http.StatusBadRequest)

			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		l.Info(string(body))
		type GenerateRequest struct {
			Width  string `json:"width"`
			Height string `json:"height"`
		}
		gr := &GenerateRequest{}
		err = json.Unmarshal(body, gr)
		if err != nil {
			l.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if gr.Width == "" || gr.Height == "" {
			err = errors.New("empty width or height")
			l.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		width, err := strconv.Atoi(gr.Width)
		if err != nil {
			err = errors.Wrap(err, "wrong width format")
			l.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}
		height, err := strconv.Atoi(gr.Height)
		if err != nil {
			err = errors.Wrap(err, "wrong height format")
			l.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		g.Width = width
		g.Height = height
		err = g.Generate()
		if err != nil {
			err = errors.Wrap(err, "error generating grid")
			l.Error(err)
		}
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	s := &http.Server{
		Addr:           ":80",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	l.Infof("starting endpoint on port :80")
	l.Fatal(s.ListenAndServe())
}
