package main

import (
	"fmt"
	client "github.com/github-exporter/client"
	"github.com/github-exporter/seats"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	newClient := client.NewClient()
	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "github",
			Name:      "free_seats",
			Help:      "Number of free seats available.",
		},
		func() float64 {

			freeseats := seats.GetFreeSeats(newClient)
			return float64(freeseats)
		})); err == nil {
		fmt.Println("Metric free_seats registered.")
	}
	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "github",
			Name:      "total_seats",
			Help:      "Number of total seats.",
		},
		func() float64 {

			totalseats := seats.GetTotalSeats(newClient)
			return float64(totalseats)
		})); err == nil {
		fmt.Println("Metric total_seats registered.")
	}
	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "github",
			Name:      "filled_seats",
			Help:      "Number of filled seats.",
		},
		func() float64 {

			filledseats := seats.GetFilledSeats(newClient)
			return float64(filledseats)
		})); err == nil {
		fmt.Println("Metric filled_seats registered.")
	}

	//recordMetrics()
	fmt.Println("Running on port 2112")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
