// Package trace provides initialization of the Jaeger and sending
// requests to the server
package trace

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

var (
	ErrNoHost = errors.New("host is not defined")
)

// Trace provides supporting of jaeger metrics
type Trace struct {
	host         string
	batchSize    int
	sendInterval time.Duration
	sendTimeout  time.Duration
	client       http.Client
	batch        []*trace.SpanData
	mu           sync.Mutex
	timer        *time.Timer
	exporter     *jaeger.Exporter
}

// New creates Trace instance
func New(host string) (*Trace, error) {
	if host == "" {
		return nil, ErrNoHost
	}

	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          2,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: host,
		Process: jaeger.Process{
			ServiceName: "go-blog",
			Tags: []jaeger.Tag{
				jaeger.StringTag("hostname", "localhost"),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create Jaeger exporter: %v", err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	e := Trace{
		host: host,
		client: http.Client{
			Transport: &tr,
		},
		exporter: exporter,
	}

	return &e, nil
}
