package internal

import (
	"fmt"
	"net/http"
)

func GetReport(w http.ResponseWriter, r *http.Request) {
	report := fmt.Sprintf("Plays: %v, Pauses: %v, Stops: %v", plays, pauses, stops)
	w.Write([]byte(report))
}
