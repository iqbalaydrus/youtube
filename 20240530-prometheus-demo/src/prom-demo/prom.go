package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UserRegistered = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "promdemo_user_registered_total",
		},
	)
	UserReadArticle = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "promdemo_user_read_article_total",
		},
		[]string{"user_id"},
	)
	RequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "promdemo_request_total",
		},
		[]string{"response_status"},
	)
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "promdemo_request_duration_seconds",
			Buckets: []float64{.00005, .0005, .005, .01, .025, .05, .1, .25, .5, 1, 2.5},
		},
		[]string{"response_status"},
	)
)
