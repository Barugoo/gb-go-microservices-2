package main

import (
	"time"
)

type Tracker struct {
	service, query string
	t              time.Time
	ok, status     string
}

func NewTracker(service, query, ok string) *Tracker {
	t := &Tracker{service, query, time.Now(), ok, ok}
	RequestsTotal.WithLabelValues(t.service, t.query).Inc()
	return t
}

func (t *Tracker) Done() {
	if t.status != t.ok {
		ErrorsTotal.WithLabelValues(t.service, t.query, t.status).Inc()
	}
	Duration.WithLabelValues(t.service, t.query, t.status).Observe(time.Since(t.t).Seconds())
}

func (t *Tracker) SetStatus(s string) {
	t.status = s
}
