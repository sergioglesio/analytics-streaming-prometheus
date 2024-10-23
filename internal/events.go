package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Estrutura de um evento
type Event struct {
	UserID    string `json:"user_id"`
	VideoID   string `json:"video_id"`
	EventType string `json:"event_type"` // play, pause, stop
	Timestamp int64  `json:"timestamp"`
}

var (
	plays = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "video_plays_total",
		Help: "Total number of video plays",
	})
	pauses = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "video_pauses_total",
		Help: "Total number of video pauses",
	})
	stops = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "video_stops_total",
		Help: "Total number of video stops",
	})
)

func init() {
	prometheus.MustRegister(plays, pauses, stops)
}

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Event received: %v", event)

	switch event.EventType {
	case "play":
		plays.Inc()
	case "pause":
		pauses.Inc()
	case "stop":
		stops.Inc()
	}

	w.WriteHeader(http.StatusOK)
}
