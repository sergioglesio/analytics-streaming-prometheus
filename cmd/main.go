package main

import (
	"analytics-streaming/internal"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sergioglesio/analytics-streaming-prometheus/internal"
)

func main() {
	// Endpoints para coletar eventos e expor métricas
	http.HandleFunc("/event", internal.HandleEvent)           // Endpoint para registrar eventos (play, pause, stop)
	http.HandleFunc("/report", internal.GetReport)            // Endpoint para relatórios gerais
	http.HandleFunc("/report/video", internal.GetVideoReport) // Relatório detalhado por vídeo
	http.Handle("/metrics", promhttp.Handler())               // Endpoint do Prometheus para métricas

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
