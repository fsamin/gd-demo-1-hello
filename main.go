// gd-demo-1-hello is step 1 of the git-deploy demo: the smallest repository
// that deploys. A Dockerfile with EXPOSE is the whole contract — nothing about
// the platform is committed here.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// version is bumped on each demo commit to make the rolling update visible.
const version = "v3"

func main() {
	log.Printf("gd-demo-1-hello %s starting on :8080", version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		hostname, _ := os.Hostname()
		printf(w, "Hello from gd-demo-1-hello %s\n", version)
		printf(w, "\npod:     %s\n", hostname)
		printf(w, "served:  %s\n", time.Now().Format(time.RFC3339))

		printf(w, "\nStep 1 of the git-deploy demo: a repository, a Dockerfile, a URL.\n")
		printf(w, "Nothing else is committed here — no manifest, no pipeline, no YAML.\n")
		printf(w, "Bump `version` in main.go, push, and this page changes on its own.\n")

		log.Printf("%s %s from %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Round(time.Millisecond))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// printf writes one line of the page, discarding the write error on purpose: a
// client hanging up mid-response is this server's normal weather, and there is
// nothing left to send it. Discarding it explicitly keeps the quality gate green
// if this repository ever declares one.
func printf(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(w, format, args...)
}
