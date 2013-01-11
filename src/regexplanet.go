package regexplanet

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func write_with_callback(w http.ResponseWriter, callback string, v interface{}) {

	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Max-Age", "604800") // 1 week

	var b []byte
	var err error
	b, err = json.Marshal(v)
	if err != nil {
		b = []byte("{\"success\":false,\"html\":\"<p>json.Marshal failed</p>\"}")
	}

	if callback > "" {
		w.Write([]byte(callback))
		w.Write([]byte("("))
		w.Write(b)
		w.Write([]byte(");"))
	} else {
		w.Write(b)
	}
}

type Status struct {
	Success  bool		`json:"success"`
	Hostname string
	Getwd    string
	TempDir  string
	Environ  []string
	Version  string
	Seconds  int64
}

func status_handler(w http.ResponseWriter, r *http.Request) {
	var err error
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
	status.Environ = os.Environ()
	status.Version = runtime.Version()
	status.Seconds = time.Now().Unix()
	status.Success = true

	write_with_callback(w, r.FormValue("callback"), status)
}

type TestResult struct {
	Success bool		`json:"success"`
	Html    string		`json:"html"`
	Message	string		`json:"message,omitempty"`
}

func test_handler(w http.ResponseWriter, r *http.Request) {

	retVal := TestResult{}

	retVal.Success = true
	retVal.Html = "<div class=\"alert alert-warning\">Actually, it is a lot less than beta: the real code isn't even written yet!</div>"

	write_with_callback(w, r.FormValue("callback"), retVal)
}
