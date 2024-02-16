package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type BotMessage struct {
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user"`

	Room struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"room"`

	Message struct {
		ID   int `json:"id"`
		Body struct {
			HTML  string `json:"html"`
			Plain string `json:"plain"`
		} `json:"body"`
	} `json:"message"`
}

func main() {
	http.HandleFunc("/trace", traceHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	panic(http.ListenAndServe(":"+port, nil))
}

func traceHandler(w http.ResponseWriter, r *http.Request) {
	var msg BotMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	uri, err := url.Parse(msg.Message.Body.Plain)
	if err != nil || (uri.Scheme != "http" && uri.Scheme != "https") {
		fmt.Fprintln(w, "That doesn't look like a valid URL for to me to call")
		return
	}

	response, err := timeRequest(msg.User.Name, uri)
	if err != nil {
		fmt.Fprintf(w, "Failed to time the request (%s)", err)
		return
	}

	fmt.Fprintln(w, response)
}

func timeRequest(username string, uri *url.URL) (string, error) {
	trace, err := TraceRequest(uri)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Hi %s! I've checked %s for you, and here's what I found:<br><br>\n", username, uri)
	result += formatDuration("DNS lookup", trace.DNSStart, trace.DNSDone)
	result += formatDuration("Connect", trace.ConnectStart, trace.ConnectDone)
	result += formatDuration("TLS negotiation", trace.TLSStart, trace.TLSDone)
	result += formatDuration("Sending headers", trace.ConnectionReady(), trace.WroteHeaders)
	result += formatDuration("Time to first byte", trace.WroteHeaders, trace.FirstByte)
	result += formatDuration("Time to last byte", trace.WroteHeaders, trace.AllDone)

	return result, nil
}

func errorResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %s", msg)
}

func formatDuration(label string, start, end time.Time) string {
	dur := end.Sub(start).String()
	if start.IsZero() || end.IsZero() {
		dur = "n/a"
	}
	return fmt.Sprintf("%s: <strong>%s</strong><br>\n", label, dur)
}
