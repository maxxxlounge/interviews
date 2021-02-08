package main

import (
	"encoding/json"
	"github.com/maxxxlounge/interviews/GameOfLife/game"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"github.com/pkg/errors"
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
		if r.Method != http.MethodGet{
			http.Error(w,"can't get resource using not GET method", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		out, err := json.Marshal(g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(out)
	})

	http.HandleFunc("/cells/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w,"can't change resorce using not POST method", http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		l.Info(string(body))
		type GenerateRequest struct {
			Width string `json:"width"`
			Height string `json:"height"`
		}
		gr := &GenerateRequest{}
		err = json.Unmarshal(body,gr)
		if err != nil {
			l.Error(err)
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		width,err := strconv.Atoi(gr.Width)
		if err != nil {
			errors.Wrap(err,"wrong width format")
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		height,err := strconv.Atoi(gr.Height)
		if err != nil {
			errors.Wrap(err,"wrong height format")
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}

		g.Width = width
		g.Height = height
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

