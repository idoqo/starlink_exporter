// Updated using Starlink Dish Firmware Version - 29424243-0ba5-4e9b-b402-79d25cb6f8de.uterm.release

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"github.com/idoqo/starlink_exporter/internal/exporter"
)

const (
	metricsPath = "/metrics"
)

func main() {
	port := flag.String("port", "9817", "listening port to expose metrics on")
	dishAddress := flag.String("dish-address", exporter.DishAddress, "IP address and port to reach dish")
	routerAddress := flag.String("router-address", exporter.RouterAddress, "IP address and port to reach router")
	flag.Parse()

	exp, err := exporter.New(*dishAddress, *routerAddress)
	if err != nil {
		log.Fatalf("could not start exporter: %s", err.Error())
	}
	defer exp.Conn.Close()
	log.Infof("dish id: %s", exp.DishID)
	log.Infof("router id: %s", exp.RouterID)

	r := prometheus.NewRegistry()
	r.MustRegister(exp)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>Starlink Exporter</title></head>
             <body>
             <h1>Starlink Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             <p><a href='/health'>Health (gRPC connection state to Starlink dish)</a></p>
             </body>
             </html>`))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		switch exp.Conn.GetState() {
		case 0, 2:
			// Idle or Ready
			w.WriteHeader(http.StatusOK)
		case 1, 3:
			// Connecting or TransientFailure
			w.WriteHeader(http.StatusServiceUnavailable)
		case 4:
			// Shutdown
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, _ = fmt.Fprintf(w, "%s\n", exp.Conn.GetState())
	})

	http.Handle(metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
