package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	IncSuccessCreateTrackCounter()
	IncSuccessRemoveTrackCounter()
	IncSuccessUpdateTrackCounter()
	IncGetTrackCounter()
}

type metrics struct {
	successCreateTrackCounter prometheus.Counter
	successRemoveTrackCounter prometheus.Counter
	successUpdateTrackCounter prometheus.Counter
	successGetTrackCounter    prometheus.Counter
}

func NewApiMetrics() Metrics {
	return &metrics{
		successCreateTrackCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_create_track_request_count",
			Help: "The total number of success created tracks",
		}),
		successRemoveTrackCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_remove_track_request_count",
			Help: "The total number of success removed tracks",
		}),
		successUpdateTrackCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_update_track_request_count",
			Help: "The total number of success updated tracks",
		}),
		successGetTrackCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "success_get_track_request_count",
			Help: "The total number of success updated tracks",
		}),
	}
}

func (m *metrics) IncSuccessCreateTrackCounter() {
	m.successCreateTrackCounter.Inc()
}
func (m *metrics) IncSuccessRemoveTrackCounter() {
	m.successRemoveTrackCounter.Inc()
}

func (m *metrics) IncSuccessUpdateTrackCounter() {
	m.successUpdateTrackCounter.Inc()
}

func (m *metrics) IncGetTrackCounter() {
	m.successGetTrackCounter.Inc()
}
