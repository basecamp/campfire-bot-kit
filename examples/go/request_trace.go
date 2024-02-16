package main

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"
)

type RequestTrace struct {
	Started      time.Time
	DNSStart     time.Time
	DNSDone      time.Time
	ConnectStart time.Time
	ConnectDone  time.Time
	TLSStart     time.Time
	TLSDone      time.Time
	WroteHeaders time.Time
	FirstByte    time.Time
	AllDone      time.Time
}

func (t RequestTrace) ConnectionReady() time.Time {
	if t.TLSDone.After(t.ConnectDone) {
		return t.TLSDone
	}
	return t.ConnectDone
}

func TraceRequest(uri *url.URL) (RequestTrace, error) {
	var trace RequestTrace

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ctx = httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		GotConn:              func(info httptrace.GotConnInfo) { trace.Started = time.Now() },
		DNSStart:             func(info httptrace.DNSStartInfo) { trace.DNSStart = time.Now() },
		DNSDone:              func(info httptrace.DNSDoneInfo) { trace.DNSDone = time.Now() },
		ConnectStart:         func(network, addr string) { trace.ConnectStart = time.Now() },
		ConnectDone:          func(network, addr string, err error) { trace.ConnectDone = time.Now() },
		TLSHandshakeStart:    func() { trace.TLSStart = time.Now() },
		TLSHandshakeDone:     func(tls.ConnectionState, error) { trace.TLSDone = time.Now() },
		WroteHeaders:         func() { trace.WroteHeaders = time.Now() },
		GotFirstResponseByte: func() { trace.FirstByte = time.Now() },
	})

	req, err := http.NewRequest("GET", uri.String(), nil)
	req = req.WithContext(ctx)

	if err != nil {
		return trace, err
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableKeepAlives = true
	c := &http.Client{Transport: t}

	resp, err := c.Do(req)
	if err != nil {
		return trace, err
	}
	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)
	trace.AllDone = time.Now()

	return trace, nil
}
