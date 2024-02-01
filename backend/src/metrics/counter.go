package metrics

import "github.com/prometheus/client_golang/prometheus"

var DatabaseCallCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "estate_market_database_calls_total",
		Help: "Number of database calls",
	}, []string{"type_name", "operation_name", "status"},
)
