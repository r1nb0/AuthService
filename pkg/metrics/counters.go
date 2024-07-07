package metrics

import "github.com/prometheus/client_golang/prometheus"

var DBCall = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "db_call_count",
		Help: "Number of database calls",
	}, []string{"type_name", "operation_name", "status"},
)
