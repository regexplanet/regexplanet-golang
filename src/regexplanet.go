package regexplanet

import (
	"fmt"
	"http"
	"json"
	"os"
	"runtime"
	"time"
)

func init() {
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/status.json", status_handler)
	http.HandleFunc("/test.json", test_handler)
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, RegexPlanet!")
}

type Status struct {
	Success bool
	Hostname string
	Getwd string
	TempDir string
	Envs []string
	Version string
	Seconds int64
}

func status_handler(w http.ResponseWriter, r *http.Request) {
	var err os.Error
	status := Status{}

	status.Getwd, err = os.Getwd()
	if err != nil {
		status.Getwd = "ERROR!"
	}

	status.Hostname, err = os.Hostname()
	if err != nil {
		status.Hostname = "ERROR"
	}

	status.TempDir = os.TempDir()
	status.Envs = os.Envs
	status.Version = runtime.Version()
	status.Seconds = time.Seconds()
	status.Success = true

	var b []byte
	b, err = json.Marshal(status)
	if err != nil {
		return
	}

	if b[2] == 'S' {		// HACK: it doesn't get much hackier than this, but json.Marshal doesn't marshal lower-case members.  Is there a way around this?
		b[2] = 's'
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Write(b)
}

type TestResult struct {
	Success bool
	Html string
}

func test_handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Max-Age", "604800") 		// 1 week

	retVal := TestResult{}

	retVal.Success = true
	retVal.Html = "<div class=\"alert alert-warning\">Actually, it is a lot less than beta: the real code isn't even written yet!</div>"

	var err os.Error
	var b []byte
	b, err = json.Marshal(retVal)
	if err != nil {
		fmt.Fprint(w, "{\"success\":false,\"html\":\"<p>json.Marshal failed</p>\"}")
		return
	}

	if b[2] == 'S' {		// HACK: it doesn't get much hackier than this, but json.Marshal doesn't marshal lower-case members.  Is there a way around this?
		b[2] = 's'
	}

	if b[17] == 'H' {
		b[17] = 'h'
	}

	w.Write(b)
}